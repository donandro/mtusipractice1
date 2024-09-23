package grpc

import (
	"context"
	"fmt"
	"net"

	grpcconfig "pillsreminder/apps/pillsservice/api/grpc/config"
	"pillsreminder/apps/pillsservice/api/grpc/handler"
	"pillsreminder/apps/pillsservice/config"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	gRPCServer *grpc.Server
	config     grpcconfig.GrpcConfig
	logger     zerolog.Logger
}

func NewGrpcServer(ctx context.Context, config config.AppConfig, logger zerolog.Logger) *GrpcServer {
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Error().Err(err).Msg("Recovered from panic")

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))

	handler.Register(gRPCServer, logger)

	return &GrpcServer{
		gRPCServer: gRPCServer,
		config:     config.Grpc,
		logger:     logger,
	}
}

func (a *GrpcServer) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.Server.Port))
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to run gRPC server")
		panic("failed to run gRPC server")
	}

	a.logger.Info().Str("address", listener.Addr().String()).Msg("gRPC server started")

	err = a.gRPCServer.Serve(listener)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to handle gRPC server messages")
		panic("failed to handle gRPC server messages")
	}
}

func (a *GrpcServer) GracefulStop() {
	a.logger.Info().Msg("grpc server graceful stop")
	a.gRPCServer.GracefulStop()
}
