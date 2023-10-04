package router

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/question/controller"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddQuestionRouter(db *gorm.DB, r *gin.RouterGroup) {
	qc := getQuestionController(db)

	r.POST("/questions", qc.CreateQuestion)
}

func getQuestionController(db *gorm.DB) controller.IQuestionController {
	questionRepo := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepo)
	questionController := controller.NewCommentController(questionService)
	return questionController
}
