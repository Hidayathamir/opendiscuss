package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/user/repository"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req dto.ReqRegisterUser) (int, error)
	LoginUser(ctx context.Context, req dto.ReqLoginUser) (string, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}
