package usecases

import (
	"context"

	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/pkg/http_utils"
)

type UserUsecase struct {
	userDBservice UserDBService
}

func NewUserCase(userDBservice UserDBService) *UserUsecase {
	return &UserUsecase{userDBservice: userDBservice}
}

func (userUC *UserUsecase) CreateUser(ctx context.Context, userEntity *entities.UserEntity) (*entities.UserEntity, *http_utils.ErrorResponse) {
	err := userUC.userDBservice.Store(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func (userUC *UserUsecase) UpdatePassword(ctx context.Context, password string) *http_utils.ErrorResponse {
	return nil
}

func (userUC *UserUsecase) GetUser(ctx context.Context, userID string) (*entities.UserEntity, *http_utils.ErrorResponse) {
	return userUC.userDBservice.Get(ctx, userID)
}
