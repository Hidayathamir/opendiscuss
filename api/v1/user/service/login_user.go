package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/pkg/errors"
)

func (us *UserService) LoginUser(ctx context.Context, req dto.ReqLoginUser) (string, error) {
	if err := req.Validate(); err != nil {
		return "", errors.Wrap(err, constant.ERR_REQ_BODY_VALIDATE)
	}

	user, err := us.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return "", errors.Wrap(err, "error get user by username")
	}

	err = utils.CompareHashAndPassword(user.Password, req.Password)
	if err != nil {
		return "", errors.Wrap(err, "wrong password")
	}

	tokenString, err := utils.GenerateUserJWTToken(user.ID)
	if err != nil {
		return "", errors.Wrap(err, "error generate user jwt token")
	}

	return tokenString, err
}
