package controllers

import (
	"pillsreminder/apps/mobileapi/api/rest/models"
	"pillsreminder/internal/mobileapi/businesslayer"
	blModels "pillsreminder/internal/mobileapi/businesslayer/auth/models"
	commongin "pillsreminder/pkg/rest/gin"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthController struct {
	authProcessor businesslayer.IAuthProcessor
	logger        zerolog.Logger
}

func NewAuthController(authProcessor businesslayer.IAuthProcessor, logger zerolog.Logger) *AuthController {
	return &AuthController{
		authProcessor: authProcessor,
		logger:        logger,
	}
}

// SignUp godoc
//
//	@Summary		SignUp
//	@Description	Creates account. Generate access token.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignUpRequest	true	"Request"
//	@Success		200 {object}	models.AuthResponse
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/signup [post]
func (ac *AuthController) SignUp(c *gin.Context) {
	var request models.SignUpRequest
	err := c.BindJSON(&request)
	if err != nil {
		commongin.BadRequest(c, "Invalid request")
		return
	}

	result, err := ac.authProcessor.SignUp(
		c.Request.Context(), blModels.User{
			Name:     *request.Name,
			Login:    *request.Login,
			Password: *request.Password,
		})
	if err != nil {
		commongin.InternalServerError(c)
		return
	}

	// not authorized - user exists
	if result == nil {
		commongin.Conflict(c)
		return
	}

	commongin.Ok(c, models.AuthResponse{
		AccessToken: result.AccessToken,
		Expiry:      result.Expiry,
	})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Authorization by user and password
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LoginRequest	true	"Request"
//	@Success		200 {object}	models.AuthResponse
//	@Failure		500	{object}	models.MessageResponse
//	@Router			/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var request models.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		commongin.BadRequest(c, "Invalid request")
		return
	}

	result, err := ac.authProcessor.Login(c.Request.Context(), *request.Login, *request.Password)
	if err != nil {
		commongin.InternalServerError(c)
		return
	}

	// not authorized - password wrong or user dows not exists
	if result == nil {
		commongin.Conflict(c)
		return
	}

	commongin.Ok(c, models.AuthResponse{
		AccessToken: result.AccessToken,
		Expiry:      result.Expiry,
	})
}
