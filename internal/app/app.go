package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/interceptor"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_service"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	*grpc_service.Implementation
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp(service *domain.InquiryService) *App {
	a := &App{}
	a.initDeps(service)
	return a
}

func (a *App) initDeps(service *domain.InquiryService) {
	a.initGRPCServer(service)
}

func (a *App) initGRPCServer(service *domain.InquiryService) {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptor.LoggingInterceptor),
		),
	)

	reflection.Register(s)

	inquiry.RegisterInquiryServer(s, grpc_service.NewInquiryV1(service))
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
