package controller

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/answer/service"
	"github.com/gin-gonic/gin"
)

type IAnswerController interface {
	CreateAnswer(ctx *gin.Context)
	GetAnswerListByQuestionID(ctx *gin.Context)
	GetAnswerByID(ctx *gin.Context)
}

type AnswerController struct {
	service service.IAnswerService
}

func NewAnswerController(service service.IAnswerService) IAnswerController {
	return &AnswerController{service: service}
}
