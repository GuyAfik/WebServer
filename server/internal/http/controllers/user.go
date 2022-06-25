package controllers

import (
	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/internal/usecases"
	"net/http"

	"github.com/WebServer/server/pkg/http_utils"
	"github.com/WebServer/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	userService usecases.UserService
	logger      logger.LoggerService
}

func newUserRoutes(userService usecases.UserService, logger logger.LoggerService) *userRoutes {
	return &userRoutes{userService: userService, logger: logger}
}

func createUserRoutes(handler *gin.RouterGroup, userService usecases.UserService, logger logger.LoggerService) {
	userRoutes := newUserRoutes(userService, logger)
	userGroup := handler.Group("/users")
	{
		userGroup.GET("/:id", userRoutes.getUser) // get single user
		userGroup.POST("", userRoutes.createUser) // create single user
		userGroup.PUT("/:id")                     // update password
		userGroup.GET("", userRoutes.getAllUsers) // get all users
	}
}

// @Summary     Show a single user
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     404 not found or 500 internal error
// @Router      /users/get/:id/
func (userRoutes *userRoutes) getUser(c *gin.Context) {
	user, err := userRoutes.userService.GetUser(c.Request.Context(), c.Param("id"))
	if err != nil {
		userRoutes.logger.Error(err, "Get-User")
		err.Response(c)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary     Show all available users
// @Accept      json
// @Produce     json
// @Success     200 (empty slice in case there aren't any users)
// @Router      /users
func (userRoutes *userRoutes) getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userRoutes.userService.GetAllUsers(c))
}

// @Summary     Create a single user
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     400 bad request or 500 internal error
// @Router      /users/create
func (userRoutes *userRoutes) createUser(c *gin.Context) {
	var request entities.UserEntity
	if err := c.ShouldBindJSON(&request); err != nil {
		customErr := http_utils.NewErrorResponse(err.Error(), http.StatusBadRequest)
		userRoutes.logger.Error(err, "Create-User")
		customErr.Response(c)
		return
	}

	user, err := userRoutes.userService.CreateUser(c.Request.Context(), &request)
	if err != nil {
		userRoutes.logger.Error(err, "Create-User")
		err.Response(c)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary     Update user password
// @Accept      json
// @Produce     json
// @Success     204
// @Failure     404 not found or 500 internal error
// @Router      /users/update/:id/password
func (userRoutes *userRoutes) updatePassword(c *gin.Context) {
	var request entities.UserEntity
	if err := c.ShouldBindJSON(&request); err != nil {
		errResponse := &http_utils.ErrorResponse{Message: err.Error(), Code: http.StatusBadRequest}
		userRoutes.logger.Error(errResponse, "updatePassword")
		errResponse.Response(c)
		return
	}

	err := userRoutes.userService.UpdatePassword(c.Request.Context(), request.Password)
	if err != nil {
		userRoutes.logger.Error(err, "updatePassword")
		err.Response(c)
		return
	}

	c.JSON(http.StatusNoContent, "")
}
