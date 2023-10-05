package controller

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/user/service"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) IUserController {
	return &UserController{service: service}
}
