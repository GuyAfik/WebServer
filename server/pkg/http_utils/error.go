package http_utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{Message: message, Code: code}
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("Error: %s, code: %d", err.Message, err.Code)
}

func (err *ErrorResponse) Response(c *gin.Context) {
	c.AbortWithStatusJSON(err.Code, err)
}
