package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/auth"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain"
	domainModel "github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
	"github.com/romanzimoglyad/inquiry-backend/internal/interceptor"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server"
	"github.com/romanzimoglyad/inquiry-backend/internal/logger"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	service *domain.InquiryService
	*grpc_server.Implementation
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp(service *domain.InquiryService) *App {
	a := &App{service: service}
	a.initDeps(service)
	return a
}

func (a *App) initDeps(service *domain.InquiryService) {
	a.initGRPCServer(service)
}

func (a *App) initGRPCServer(service *domain.InquiryService) {
	mtr, err := service.GetMethodToRole(context.Background())
	if err != nil {
		logger.Fatal().Msgf("Error in GetMethodToRole: %v", err)
	}
	m := make(map[string]int32, len(mtr))
	for _, v := range mtr {
		m[v.Method] = v.RoleId
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptor.LoggingInterceptor, interceptor.AuthInterceptor(m)),
		),
	)

	reflection.Register(s)

	inquiry.RegisterInquiryServer(s, grpc_server.NewInquiryV1(service))
	a.grpcServer = s
}

func (a *App) RunHTTPServer() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%s", config.Config.GrpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Fatal().Msgf("Failed to dial server: %v", err)
	}

	gwmux := runtime.NewServeMux()
	gwmux.HandlePath("POST", "/lesson/file", a.uploadHandler)
	// Register Greeter
	err = inquiry.RegisterInquiryHandler(context.Background(), gwmux, conn)
	if err != nil {
		logger.Fatal().Msgf("Failed to register gateway: %v", err)
	}
	handler := interceptor.CorsMiddleware(gwmux)

	a.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", config.Config.HttpPort),
		Handler:           handler,
		ReadHeaderTimeout: 100 * time.Millisecond,
	}
	logger.Info().Msgf("Serving gRPC-Gateway on: %s", config.Config.HttpPort)
	if err := a.httpServer.ListenAndServe(); err != nil {
		logger.Fatal().Err(err).Msg("error in ListenAndServe")
	}
}

func (a *App) RunGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Config.GrpcPort))
	if err != nil {
		logger.Fatal().Msgf("failed to listen: %v", err)
	}

	logger.Info().Msgf("Serving gRPC on: %s", config.Config.GrpcPort)
	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			logger.Fatal().Err(err).Msg("error in grpcServer Serve")
		}
	}()
}

func (a *App) uploadHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Parse the multipart form data
	token, ok := r.Header["Authorization"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Validate token
	claims, err := auth.ValidateToken(token[0])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !(claims.Role == int32(auth.RoleAdmin)) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // Set max memory to 10MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}
	var imgFile *domainModel.File
	// Get the img from the request
	img, fileHeader, err := r.FormFile("file")
	if err != nil {
		logger.Info().Msg("Failed to get img from form")
	} else {
		defer img.Close()

		// Read the img content
		fileContent, err := ioutil.ReadAll(img)
		if err != nil {
			http.Error(w, "Failed to read img content", http.StatusInternalServerError)
			return
		}
		imgFile = &domainModel.File{
			Name: fileHeader.Filename,
			Data: fileContent,
		}
	}
	// Get the JSON data from the request
	jsonData := r.FormValue("json")

	var request *domainModel.AddFileRequest
	err = json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["files"]
	domainFiles := make([]*domainModel.File, 0, len(files))

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		newFile := &domainModel.File{
			Name: fileHeader.Filename,
		}
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newFile.Data = content
		domainFiles = append(domainFiles, newFile)
	}
	oldFiles := r.MultipartForm.Value["oldFiles"]
	oldDomainFiles := make([]string, 0, len(oldFiles))

	for _, oldFile := range oldFiles {

		oldDomainFiles = append(oldDomainFiles, oldFile)
	}

	orderId, err := a.service.HandleFiles(r.Context(), &domainModel.AddFileRequest{
		UserId:   request.UserId,
		LessonId: request.LessonId,
		OldFiles: oldDomainFiles,
		Files:    domainFiles,
		Img:      imgFile,
	})
	if err != nil {
		http.Error(w, "Failed to HandleFile", http.StatusInternalServerError)
		return
	}
	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(orderId))
}
