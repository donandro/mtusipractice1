package handler

import (
	"context"

	mobileapigrpc "pillsreminder/internal/client/mobileapi/grpc"

	"github.com/rs/zerolog"

	"google.golang.org/grpc"

	"pillsreminder/internal/mobileapi/businesslayer"
)

type serverAPI struct {
	mobileapigrpc.UnimplementedMobileApiServiceServer
	userProcessor businesslayer.IUserProcessor
	logger        zerolog.Logger
}

func (s serverAPI) GetUserById(ctx context.Context, request *mobileapigrpc.GetUserByIdRequest) (*mobileapigrpc.GetUserByIdResponse, error) {
	user, err := s.userProcessor.GetUser(ctx, int(request.Id))
	if err != nil {
		s.logger.Error().Err(err).Msg("failed get user by id")
		return nil, err
	}

	return &mobileapigrpc.GetUserByIdResponse{
		User: &mobileapigrpc.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func Register(gRPCServer *grpc.Server, userProcessor businesslayer.IUserProcessor, logger zerolog.Logger) {
	mobileapigrpc.RegisterMobileApiServiceServer(gRPCServer, &serverAPI{
		userProcessor: userProcessor,
		logger:        logger,
	})
}
