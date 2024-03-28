package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/interceptor"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/romanzimoglyad/inquiry-backend/internal/service"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	*service.Implementation
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp() *App {
	a := &App{}
	a.initDeps()
	return a
}

func (a *App) initDeps() {
	a.initGRPCServer()
}

func (a *App) initGRPCServer() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(interceptor.LoggingInterceptor),
		),
	)

	reflection.Register(s)

	inquiry.RegisterInquiryServer(s, service.NewInquiryV1())
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

	a.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", config.Config.HttpPort),
		Handler:           gwmux,
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
