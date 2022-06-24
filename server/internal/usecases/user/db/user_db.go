package db

import (
	"context"

	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/pkg/http_utils"
)

type UserDBService struct {
}

func New() *UserDBService {
	return &UserDBService{}
}

func (uDBs *UserDBService) Get(ctx context.Context, userID string) (*entities.UserEntity, *http_utils.ErrorResponse) {
	return &entities.UserEntity{Username: "1234", Password: "1234", Email: "test@gmail.com"}, nil
}

func (uDBs *UserDBService) Store(ctx context.Context, userEntity *entities.UserEntity) *http_utils.ErrorResponse {
	return nil
}
