package router

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/question/controller"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/service"
	"github.com/Hidayathamir/opendiscuss/middleware"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddQuestionRouter(db *gorm.DB, r *gin.RouterGroup) {
	qc := getQuestionController(db)

	r.POST("/questions", middleware.Authenticate, qc.CreateQuestion)
	r.GET("/questions", qc.GetQuestionList)
}

func getQuestionController(db *gorm.DB) controller.IQuestionController {
	questionRepo := repository.NewQuestionRepository(db)
	trManager := utils.NewTransactionManager(db)
	questionService := service.NewQuestionService(questionRepo, trManager)
	questionController := controller.NewQuestionController(questionService)
	return questionController
}
