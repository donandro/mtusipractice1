package auth

import (
	"context"
	"errors"
	"time"

	"pillsreminder/internal/mobileapi/businesslayer/auth/claims"
	"pillsreminder/internal/mobileapi/businesslayer/auth/config"
	"pillsreminder/internal/mobileapi/businesslayer/auth/models"
	"pillsreminder/internal/mobileapi/datalayer"
	dlmodels "pillsreminder/internal/mobileapi/datalayer/postgres/models"
	"pillsreminder/pkg/jwt"

	"github.com/rs/zerolog"
)

type AuthProcessor struct {
	userRepo     datalayer.IUserRepo
	secureConfig config.Secure
	logger       zerolog.Logger
}

func NewAuthProcessor(
	userRepo datalayer.IUserRepo,
	secureConfig config.Secure,
	logger zerolog.Logger,
) *AuthProcessor {
	return &AuthProcessor{
		userRepo:     userRepo,
		secureConfig: secureConfig,
		logger:       logger,
	}
}

func (ap *AuthProcessor) SignUp(ctx context.Context, userInfo models.User) (*models.Authentication, error) {
	existingUser, err := ap.userRepo.GetByLogin(userInfo.Login)
	if err != nil {
		return nil, errors.New("failed to get user by login")
	}

	if existingUser != nil {
		return nil, nil
	}

	dbUser, err := ap.userRepo.Add(dlmodels.UserDB{
		Name:     userInfo.Name,
		Login:    userInfo.Login,
		Password: userInfo.Password,
	})
	if err != nil {
		return nil, errors.New("failed to get add user")
	}

	auth, err := ap.generateTokens(ctx, dbUser.ID, ap.secureConfig.JwtSecret)
	if err != nil {
		ap.logger.Err(err).Ctx(ctx).Msg("failed to generate tokens")
		return nil, errors.New("failed to generate token")
	}

	return auth, nil
}

func (ap *AuthProcessor) Login(ctx context.Context, login string, password string) (*models.Authentication, error) {
	existingUser, err := ap.userRepo.GetByLogin(login)
	if err != nil {
		return nil, errors.New("failed to get user by login")
	}

	if existingUser == nil || existingUser.Password != password {
		return nil, nil
	}

	auth, err := ap.generateTokens(ctx, existingUser.ID, ap.secureConfig.JwtSecret)
	if err != nil {
		ap.logger.Err(err).Ctx(ctx).Msg("failed to generate tokens")
		return nil, errors.New("failed to generate tokens")
	}

	return auth, nil
}

func (ap *AuthProcessor) generateTokens(ctx context.Context, userID int, jwtSecret string) (*models.Authentication, error) {
	claims := map[string]interface{}{
		claims.UserID: userID,
	}

	expires := time.Now().Add(ap.secureConfig.AccessTokenLifetime)

	accessToken, err := jwt.GenerateToken(
		claims,
		jwtSecret,
		expires,
	)
	if err != nil {
		ap.logger.Err(err).Ctx(ctx).Msg("failed to generate access token")
		return nil, errors.New("failed to generate access token")
	}

	return &models.Authentication{
		AccessToken: accessToken,
		Expiry:      int(expires.Unix()),
	}, nil
}
