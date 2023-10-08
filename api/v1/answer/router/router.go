package router

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/answer/controller"
	"github.com/Hidayathamir/opendiscuss/api/v1/answer/repository"
	"github.com/Hidayathamir/opendiscuss/api/v1/answer/service"
	"github.com/Hidayathamir/opendiscuss/middleware"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAnswerRouter(db *gorm.DB, r *gin.RouterGroup) {
	ac := getAnswerController(db)

	r.POST("/questions/:questionid/answers", middleware.Authenticate, ac.CreateAnswer)
	r.GET("/questions/:questionid/answers", ac.GetAnswerListByQuestionID)
	r.GET("/answers/:answerid", ac.GetAnswerByID)
}

func getAnswerController(db *gorm.DB) controller.IAnswerController {
	AnswerRepo := repository.NewAnswerRepository(db)
	trManager := utils.NewTransactionManager(db)
	AnswerService := service.NewAnswerService(AnswerRepo, trManager)
	AnswerController := controller.NewAnswerController(AnswerService)
	return AnswerController
}
