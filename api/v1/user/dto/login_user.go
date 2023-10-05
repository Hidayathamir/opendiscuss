package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r ReqLoginUser) ToModelUser() model.User {
	return model.User{
		Username: r.Username,
		Password: r.Password,
	}
}

func (r ReqLoginUser) Validate() error {
	if r.Username == "" {
		return errors.New("username can not be empty")
	}
	if r.Password == "" {
		return errors.New("password can not be empty")
	}
	return nil
}

type ResLoginUser struct {
	Token string `json:"token"`
}
