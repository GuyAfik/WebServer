package app

import (
	"github.com/WebServer/server/config"
	"github.com/WebServer/server/internal/http/controllers"
	"github.com/WebServer/server/internal/usecases"
	"github.com/WebServer/server/internal/usecases/user/db"
	"github.com/WebServer/server/pkg/httpApi"
	"github.com/WebServer/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RunServer(config *config.Config) {
	logger := logger.New(config.Log.Level)

	userUseCase := usecases.NewUserCase(db.New())
	// HTTP Server
	handler := gin.New()
	controllers.NewRouter(handler, userUseCase, logger)
	httpApi.New(handler, logger, httpApi.Port(config.HTTP.Port))
}