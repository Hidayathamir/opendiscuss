package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/golang-jwt/jwt"
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

	tokenString, err := us.generateUserJWTToken(ctx, user.ID)
	if err != nil {
		return "", errors.Wrap(err, "error generate user jwt token")
	}

	return tokenString, err
}

func (us *UserService) generateUserJWTToken(ctx context.Context, userID int) (string, error) {
	expireIn := time.Hour * time.Duration(environtment.JWT_EXPIRE_HOUR)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": fmt.Sprintf("%d", userID),
		"exp":     time.Now().Add(expireIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(environtment.JWT_SIGN_KEY))
	if err != nil {
		return "", errors.Wrap(err, "error sign jwt token")
	}

	tokenString = "Bearer " + tokenString

	return tokenString, nil
}
