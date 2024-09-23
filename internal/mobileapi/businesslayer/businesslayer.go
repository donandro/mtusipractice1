package businesslayer

import (
	"context"

	authModels "pillsreminder/internal/mobileapi/businesslayer/auth/models"
	userModels "pillsreminder/internal/mobileapi/businesslayer/user/models"
)

type IUserProcessor interface {
	GetUser(ctx context.Context, userID int) (*userModels.User, error)
}

type IAuthProcessor interface {
	SignUp(ctx context.Context, request authModels.User) (*authModels.Authentication, error)
	Login(ctx context.Context, login string, password string) (*authModels.Authentication, error)
}
