package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
)

func (us *UserService) RegisterUser(ctx context.Context, req dto.ReqRegisterUser) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return 0, err
	}
	req.Password = hashedPassword

	userID, err := us.repo.RegisterUser(ctx, req.ToModelUser())
	if err != nil {
		return 0, err
	}

	return userID, err
}
