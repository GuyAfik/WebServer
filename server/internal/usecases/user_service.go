package usecases

import (
	"context"

	"github.com/WebServer/server/internal/entities"
)

type (
	// UserService
	UserService interface {
		CreateUser(*context.Context, *entities.User) (*entities.User, error)
		UpdatePassword(*context.Context, *entities.User) error
		GetUser(*context.Context, *entities.User) (*entities.User, error)
	}

	// UserDBService
	UserDBService interface {
		Store(*context.Context, *entities.User) error
		Get(*context.Context, *entities.User) (*entities.User, error)
	}
)