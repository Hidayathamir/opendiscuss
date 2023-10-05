package controller

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/question/service"
	"github.com/gin-gonic/gin"
)

type IQuestionController interface {
	CreateQuestion(ctx *gin.Context)
	GetQuestionList(ctx *gin.Context)
}

type QuestionController struct {
	service service.IQuestionService
}

func NewQuestionController(service service.IQuestionService) IQuestionController {
	return &QuestionController{service: service}
}
