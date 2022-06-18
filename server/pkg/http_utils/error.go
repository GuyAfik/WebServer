package http_utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiErrorResponse struct {
	Code int
	Message string
}

func (err *ApiErrorResponse) Error() string {
	return fmt.Sprintf("Error: %s, code: %d", err.Message, err.Code)
}

func ErrorRepsonse(c *gin.Context, err *ApiErrorResponse) {
	c.AbortWithStatusJSON(err.Code, err.Message)
}