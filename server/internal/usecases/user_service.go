package usecases

import (
	"context"
	"github.com/WebServer/server/pkg/http_utils"
	"github.com/WebServer/server/internal/entities"
)

type (
	// UserService
	UserService interface {
		CreateUser(ctx context.Context, userEntity *entities.User) (*entities.User, *http_utils.ApiErrorResponse)
		UpdatePassword(ctx context.Context, password string) *http_utils.ApiErrorResponse
		GetUser(ctx context.Context, userID string) (*entities.User, *http_utils.ApiErrorResponse)
	}

	// UserDBService
	UserDBService interface {
		Store(ctx context.Context, userEntity *entities.User) *http_utils.ApiErrorResponse
		Get(ctx context.Context, userID string) (*entities.User, *http_utils.ApiErrorResponse)
	}
)