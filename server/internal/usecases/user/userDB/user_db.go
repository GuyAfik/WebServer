package userDB

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"net/http"

	"github.com/WebServer/server/internal/entities"
	"github.com/WebServer/server/pkg/http_utils"
)

type UserDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (*UserDB) TableName() string {
	return "users"
}

func (uDBs *UserDB) Get(ctx context.Context, userID uint64) (*entities.UserEntity, *http_utils.ErrorResponse) {
	userEntity := &entities.UserEntity{ID: userID}
	dbContext := uDBs.db.WithContext(ctx)

	tx := dbContext.Find(userEntity)
	if tx.RowsAffected == 0 {
		return nil, http_utils.NewErrorResponse(fmt.Sprintf("User with ID %d was not found", userID), http.StatusNotFound)
	}

	return userEntity, nil
}

func (uDBs *UserDB) Store(ctx context.Context, userEntity *entities.UserEntity) *http_utils.ErrorResponse {
	dbContext := uDBs.db.WithContext(ctx)
	tx := dbContext.Create(userEntity)
	if tx.Error != nil {
		errMessage := tx.Error.Error()
		return http_utils.NewErrorResponse(errMessage, http.StatusInternalServerError)
	}
	return nil
}

func (uDBs *UserDB) GetAll(ctx context.Context) []*entities.UserEntity {
	var userEntities []*entities.UserEntity
	dbContext := uDBs.db.WithContext(ctx)
	dbContext.Find(&userEntities)
	return userEntities
}
