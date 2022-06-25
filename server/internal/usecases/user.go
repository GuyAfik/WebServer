package usecases

import (
	"context"
	"net/http"
	"strconv"

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
	primaryKey, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, http_utils.NewErrorResponse(err.Error(), http.StatusInternalServerError)
	}
	return userUC.userDBservice.Get(ctx, primaryKey)
}

func (userUC *UserUsecase) GetAllUsers(ctx context.Context) []*entities.UserEntity {
	return userUC.userDBservice.GetAll(ctx)
}
