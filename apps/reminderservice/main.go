package main

import (
	"context"
	"os"

	"pillsreminder/apps/reminderservice/api/grpc"
	"pillsreminder/apps/reminderservice/config"
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
		Str("serviceName", "reminderservice").
		Caller().
		Logger().
		Hook(logging.LogJSONLoggerEventHook).
		Hook(logging.ContextLoggerEventHook)

	db.InitDB(config.DB)

	grpcServer := grpc.NewGrpcServer(ctx, config, logger)

	defer grpcServer.GracefulStop()
	grpcServer.Run()
}
