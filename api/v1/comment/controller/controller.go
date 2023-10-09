package controller

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/service"
	"github.com/gin-gonic/gin"
)

type ICommentController interface {
	CreateComment(ctx *gin.Context)
	GetCommentList(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
}

type CommentController struct {
	service service.ICommentService
}

func NewCommentController(service service.ICommentService) ICommentController {
	return &CommentController{service: service}
}