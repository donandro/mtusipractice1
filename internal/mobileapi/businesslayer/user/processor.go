package user

import (
	"context"
	"errors"

	"pillsreminder/internal/mobileapi/businesslayer/user/models"
	"pillsreminder/internal/mobileapi/datalayer"

	"github.com/rs/zerolog"
)

type UserProcessor struct {
	userRepo datalayer.IUserRepo
	logger   zerolog.Logger
}

func NewUserProcessor(
	userRepo datalayer.IUserRepo,
	logger zerolog.Logger,
) *UserProcessor {
	return &UserProcessor{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (ap *UserProcessor) GetUser(ctx context.Context, userID int) (*models.User, error) {
	ap.logger.Info().Ctx(ctx).Int("userID", userID).Msg("get user info by id")
	user, err := ap.userRepo.GetByID(userID)
	if err != nil || user == nil {
		ap.logger.Err(err).Ctx(ctx).Msg("failed to get user from db")
		return nil, errors.New("failed to get user from db")
	}

	return &models.User{
		ID:       0,
		Name:     user.Name,
		Email:    user.Email,
		Login:    user.Login,
		Password: user.Password,
		Created:  user.Created,
		Deleted:  user.Deleted,
	}, nil
}
