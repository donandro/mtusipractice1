package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, defaultMessageResponse{Message: "Internal server error"})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, defaultMessageResponse{Message: message})
}

func Ok(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

func OkEmpty(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func Conflict(c *gin.Context) {
	c.JSON(http.StatusConflict, nil)
}

type defaultMessageResponse struct {
	Message string
}
