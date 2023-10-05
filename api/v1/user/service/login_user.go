package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/golang-jwt/jwt"
)

func (us *UserService) LoginUser(ctx context.Context, req dto.ReqLoginUser) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	user, err := us.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	err = utils.CompareHashAndPassword(user.Password, req.Password)
	if err != nil {
		return "", errors.New("wrong password")
	}

	tokenString, err := us.generateUserJWTToken(ctx, user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (us *UserService) generateUserJWTToken(ctx context.Context, userID int) (string, error) {
	expireIn := time.Hour * time.Duration(environtment.JWT_EXPIRE_HOUR)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": fmt.Sprintf("%d", userID),
		"exp":     time.Now().Add(expireIn).Unix(),
	})

	tokenString, err := token.SignedString(environtment.JWT_SIGN_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
