package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/pkg/errors"
)

func (us *UserService) RegisterUser(ctx context.Context, req dto.ReqRegisterUser) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_BODY_VALIDATE)
	}

	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return 0, errors.Wrap(err, "error generate hash password")
	}
	req.Password = hashedPassword

	userID, err := us.repo.RegisterUser(ctx, req.ToModelUser())
	if err != nil {
		return 0, errors.Wrap(err, "error register user")
	}

	return userID, err
}
