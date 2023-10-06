package router

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/user/controller"
	"github.com/Hidayathamir/opendiscuss/api/v1/user/repository"
	"github.com/Hidayathamir/opendiscuss/api/v1/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserRouter(db *gorm.DB, r *gin.RouterGroup) {
	uc := getUserController(db)

	r.POST("/register", uc.RegisterUser)
	r.POST("/login", uc.LoginUser)
}

func getUserController(db *gorm.DB) controller.IUserController {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	return userController
}
