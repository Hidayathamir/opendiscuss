package service

import (
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/repository"
	"github.com/Hidayathamir/opendiscuss/utils"
)

type ICommentService interface {
}

type CommentService struct {
	repo      repository.ICommentRepository
	trManager utils.ITransactionManager
}

func NewCommentService(repo repository.ICommentRepository, trManager utils.ITransactionManager) ICommentService {
	return &CommentService{repo: repo, trManager: trManager}
}
