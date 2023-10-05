package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
)

type IQuestionService interface {
	CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error)
	GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error)
}

type QuestionService struct {
	repo repository.IQuestionRepository
}

func NewQuestionService(repo repository.IQuestionRepository) IQuestionService {
	return &QuestionService{repo: repo}
}
