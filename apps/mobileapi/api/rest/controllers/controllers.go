package controllers

import (
	"context"

	"pillsreminder/apps/mobileapi/config"
	"pillsreminder/internal/mobileapi/businesslayer"

	"github.com/rs/zerolog"
)

type ControllerFactory struct {
	authProcessor businesslayer.IAuthProcessor
	config        config.AppConfig
	logger        zerolog.Logger
}

func NewControllerFactory(
	ctx context.Context,
	config config.AppConfig,
	logger zerolog.Logger,
) ControllerFactory {
	processorFactory := businesslayer.NewProcessorFactory(config, logger)

	return ControllerFactory{
		authProcessor: processorFactory.GetAuthProcessor(ctx),
		config:        config,
		logger:        logger,
	}
}

func (cf *ControllerFactory) GetAuthController() *AuthController {
	return NewAuthController(
		cf.authProcessor,
		cf.logger,
	)
}

func (cf *ControllerFactory) GetPillsController() *AuthController {
	return NewPillsController(
		cf.logger,
	)
}
