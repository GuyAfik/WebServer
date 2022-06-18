package controllers

import (
	"github.com/WebServer/server/internal/usecases"
	"github.com/WebServer/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, userService usecases.UserService,  logger logger.LoggerService) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	routerGroup := handler.Group("/v1")

	createUserRoutes(routerGroup, userService, logger)
}