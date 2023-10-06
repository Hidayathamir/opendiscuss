package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(ctx context.Context, user model.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}
