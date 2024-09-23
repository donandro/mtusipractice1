package middlewares

import (
	"context"
	"net/http"
	"strings"

	"pillsreminder/internal/mobileapi/businesslayer/auth/claims"
	"pillsreminder/pkg/jwt"
	"pillsreminder/pkg/rest"

	"github.com/gin-gonic/gin"
)

func AuthBase(jwtSecret string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authorizationHeader := strings.ReplaceAll(c.GetHeader(rest.XAuthorization), "Bearer ", "")
		isValid, tokenClaims, _ := jwt.GetClaims(authorizationHeader, jwtSecret)

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), claims.UserID, tokenClaims[claims.UserID]))

		c.Next()
	})
}
