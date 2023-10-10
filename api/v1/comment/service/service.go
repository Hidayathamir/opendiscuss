package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/comment/repository"
	"github.com/Hidayathamir/opendiscuss/utils"
)

type ICommentService interface {
	CreateComment(ctx context.Context, req dto.ReqCreateComment) (int, error)
	GetCommentByID(ctx context.Context, ID int) (dto.CommentHighlight, error)
	GetCommentListByAnswerID(ctx context.Context, answerID int) ([]dto.CommentHighlight, error)
}

type CommentService struct {
	repo      repository.ICommentRepository
	trManager utils.ITransactionManager
}

func NewCommentService(repo repository.ICommentRepository, trManager utils.ITransactionManager) ICommentService {
	return &CommentService{repo: repo, trManager: trManager}
}
