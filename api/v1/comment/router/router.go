package router

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/controller"
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/repository"
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/service"
	"github.com/Hidayathamir/opendiscuss/middleware"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddCommentRouter(db *gorm.DB, r *gin.RouterGroup) {
	cc := getCommentController(db)

	r.POST("/comments", middleware.Authenticate, cc.CreateComment)
}

func getCommentController(db *gorm.DB) controller.ICommentController {
	CommentRepo := repository.NewCommentRepository(db)
	trManager := utils.NewTransactionManager(db)
	CommentService := service.NewCommentService(CommentRepo, trManager)
	CommentController := controller.NewCommentController(CommentService)
	return CommentController
}
