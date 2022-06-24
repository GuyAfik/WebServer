package usecases

import (
	"context"

	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/pkg/http_utils"
)

type (
	// UserService
	UserService interface {
		CreateUser(ctx context.Context, userEntity *entities.UserEntity) (*entities.UserEntity, *http_utils.ErrorResponse)
		UpdatePassword(ctx context.Context, password string) *http_utils.ErrorResponse
		GetUser(ctx context.Context, userID string) (*entities.UserEntity, *http_utils.ErrorResponse)
	}

	// UserDBService
	UserDBService interface {
		Store(ctx context.Context, userEntity *entities.UserEntity) *http_utils.ErrorResponse
		Get(ctx context.Context, userID string) (*entities.UserEntity, *http_utils.ErrorResponse)
	}
)
