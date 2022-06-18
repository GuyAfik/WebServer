package db

import (
	"context"

	"github.com/WebServer/server/internal/entities"
)

type UserDBService struct {

}

func New() *UserDBService {
	return &UserDBService{}
}

func (uDBs *UserDBService) Get(ctx *context.Context, user *entities.User) (*entities.User, error) {
	return user, nil
}

func (uDBs *UserDBService) Store(*context.Context, *entities.User) error {
	return nil
}