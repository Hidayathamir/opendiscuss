package utils

import (
	"fmt"
	"time"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var signingMethod = jwt.SigningMethodHS256

func GenerateUserJWTToken(userID int) (string, error) {
	expireIn := time.Hour * time.Duration(environtment.JWT_EXPIRE_HOUR)
	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		constant.JWT_CLAIM_USER_ID: fmt.Sprintf("%d", userID),
		"exp":                      time.Now().Add(expireIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(environtment.JWT_SIGN_KEY))
	if err != nil {
		return "", errors.Wrap(err, "error sign jwt token")
	}

	tokenString = "Bearer " + tokenString

	return tokenString, nil
}

func ValidateUserJWTToken(tokenString string) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(environtment.JWT_SIGN_KEY), nil
	}
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		err = errors.Wrap(err, "error parse jwt")
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return jwt.MapClaims{}, errors.Wrap(err, "auth not a token")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return jwt.MapClaims{}, errors.Wrap(err, "invalid signature")
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return jwt.MapClaims{}, errors.Wrap(err, "token expired")
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return jwt.MapClaims{}, errors.Wrap(err, "token not valid yet")
		}
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwt.MapClaims{}, errors.New("token not valid")
	}

	return claims, nil
}
