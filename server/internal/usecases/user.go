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

func (userUC *UserUsecase) CreateUser(ctx context.Context, userEntity *entities.User) (*entities.User, error) {
	err := userUC.db.Store(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func (userUC *UserUsecase) UpdatePassword(ctx context.Context, password string) error {
	return nil
}

func (userUC *UserUsecase) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	return userUC.db.Get(ctx, userID)
}
