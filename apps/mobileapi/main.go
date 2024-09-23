package main

import (
	"context"
	"os"

	"pillsreminder/apps/mobileapi/api/grpc"
	"pillsreminder/apps/mobileapi/api/rest"
	"pillsreminder/apps/mobileapi/config"
	"pillsreminder/pkg/db"
	"pillsreminder/pkg/logging"

	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()
	config := config.SetupConfig()
	logger := zerolog.New(os.Stdout).With().
		Int("ProcessID", os.Getpid()).
		Timestamp().
		Logger().With().
		Str("serviceName", "mobileapi").
		Caller().
		Logger().
		Hook(logging.LogJSONLoggerEventHook).
		Hook(logging.ContextLoggerEventHook)

	db.InitDB(config.DB)

	restServer := rest.NewRestServer(ctx, config, logger)
	grpcServer := grpc.NewGrpcServer(ctx, config, logger)

	go grpcServer.Run()
	defer grpcServer.GracefulStop()

	restServer.Run()
}
