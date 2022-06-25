package app

import (
	"github.com/WebServer/server/config"
	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/internal/http/controllers"
	"github.com/WebServer/server/internal/usecases"
	"github.com/WebServer/server/internal/usecases/user/userDB"
	"github.com/WebServer/server/pkg/dbs"
	"github.com/WebServer/server/pkg/httpApi"
	"github.com/WebServer/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunServer(config *config.Config) {
	serverLogger := logger.New(config.Log.Level)

	allEntities := []interface{}{entities.UserEntity{}}

	db, err := dbs.InitSqliteDB(allEntities)
	if err != nil {
		serverLogger.Error("Could not start Sqlite DB")
	}
	userUseCase := createUserUseCase(db)
	// HTTP Server
	handler := gin.New()
	controllers.NewRouter(handler, userUseCase, serverLogger)
	httpApi.New(handler, serverLogger, httpApi.Port(config.HTTP.Port))
}

func createUserUseCase(db *gorm.DB) *usecases.UserUsecase {
	return usecases.NewUserCase(userDB.New(db))
}
