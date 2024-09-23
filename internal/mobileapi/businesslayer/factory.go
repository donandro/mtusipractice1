package businesslayer

import (
	"context"

	"pillsreminder/apps/mobileapi/config"
	"pillsreminder/internal/mobileapi/businesslayer/auth"
	"pillsreminder/internal/mobileapi/businesslayer/user"
	"pillsreminder/internal/mobileapi/datalayer"

	"github.com/rs/zerolog"
)

type ProcessorFactory struct {
	config         config.AppConfig
	storageFactory datalayer.StorageFactory
	logger         zerolog.Logger
}

func NewProcessorFactory(
	config config.AppConfig,
	logger zerolog.Logger,
) *ProcessorFactory {
	storageFactory := datalayer.NewStorageFactory()

	return &ProcessorFactory{
		config:         config,
		storageFactory: *storageFactory,
		logger:         logger,
	}
}

func (f *ProcessorFactory) GetAuthProcessor(ctx context.Context) IAuthProcessor {
	return auth.NewAuthProcessor(
		f.storageFactory.GetUserRepo(),
		f.config.Server.Secure,
		f.logger,
	)
}

func (f *ProcessorFactory) GetUserProcessor() IUserProcessor {
	return user.NewUserProcessor(
		f.storageFactory.GetUserRepo(),
		f.logger,
	)
}
