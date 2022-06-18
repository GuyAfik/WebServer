package usecases

import (
	"context"

	"github.com/WebServer/server/internal/entities"
)

type UserUsecase struct {
	db UserDBService
}

func New(userDBservice UserDBService) *UserUsecase {
	return &UserUsecase{db: userDBservice}
}

func (userUC *UserUsecase) CreateUser(ctx *context.Context, user *entities.User) (*entities.User, error) {
	err := userUC.db.Store(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userUC *UserUsecase) UpdatePassword(ctx *context.Context, user *entities.User) error {
	return nil
}

func (userUC *UserUsecase) GetUser(ctx *context.Context, user *entities.User) (*entities.User, error) {
	return userUC.db.Get(ctx, user)
}
