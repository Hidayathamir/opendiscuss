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

	r.POST("/answers/:answerid/comments", middleware.Authenticate, cc.CreateComment)
	r.GET("/answers/:answerid/comments", cc.GetCommentListByAnswerID)
	r.GET("/comments/:commentid", cc.GetCommentByID)
	r.POST("/comments/:commentid/thumbsup", middleware.Authenticate, cc.VoteThumbsUp)
	r.POST("/comments/:commentid/thumbsdown", middleware.Authenticate, cc.VoteThumbsDown)
	r.PUT("/comments/:commentid", middleware.Authenticate, cc.UpdateCommentByID)
	r.DELETE("/comments/:commentid", middleware.Authenticate, cc.DeleteCommentByID)
}

func getCommentController(db *gorm.DB) controller.ICommentController {
	CommentRepo := repository.NewCommentRepository(db)
	trManager := utils.NewTransactionManager(db)
	CommentService := service.NewCommentService(CommentRepo, trManager)
	CommentController := controller.NewCommentController(CommentService)
	return CommentController
}
