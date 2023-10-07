package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/utils"
)

type IQuestionService interface {
	CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error)
	GetQuestionByID(ctx context.Context, ID int) (dto.QuestionHighlight, error)
	GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error)
	VoteThumbsUp(ctx context.Context, req dto.ReqVoteThumbs) (int, error)
	VoteThumbsDown(ctx context.Context, req dto.ReqVoteThumbs) (int, error)
}

type QuestionService struct {
	repo      repository.IQuestionRepository
	trManager utils.ITransactionManager
}

func NewQuestionService(repo repository.IQuestionRepository, trManager utils.ITransactionManager) IQuestionService {
	return &QuestionService{repo: repo, trManager: trManager}
}
