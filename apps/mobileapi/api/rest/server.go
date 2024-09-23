package rest

import (
	"context"

	"pillsreminder/apps/mobileapi/api/rest/controllers"
	"pillsreminder/apps/mobileapi/api/rest/router"
	"pillsreminder/apps/mobileapi/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type RestServer struct {
	router *gin.Engine
	config config.ServerConfiguration
	logger zerolog.Logger
}

func NewRestServer(ctx context.Context, config config.AppConfig, logger zerolog.Logger) *RestServer {
	controllerFactory := controllers.NewControllerFactory(ctx, config, logger)
	router := router.Setup(ctx, config, controllerFactory)

	return &RestServer{
		router: router,
		config: config.Server,
		logger: logger,
	}
}

func (a *RestServer) Run() {
	a.logger.Info().Msg("Running on port " + a.config.Port)

	err := a.router.Run(":" + a.config.Port)
	if err != nil {
		a.logger.Error().Err(err)
	}
}
