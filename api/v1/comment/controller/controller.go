package controller

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/service"
)

type ICommentController interface {
}

type CommentController struct {
	service service.ICommentService
}

func NewCommentController(service service.ICommentService) ICommentController {
	return &CommentController{service: service}
}
