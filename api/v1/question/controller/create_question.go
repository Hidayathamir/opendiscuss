package controller

import "github.com/gin-gonic/gin"

func (qc *QuestionController) CreateQuestion(ctx *gin.Context) {
	qc.service.CreateQuestion(ctx)
}
