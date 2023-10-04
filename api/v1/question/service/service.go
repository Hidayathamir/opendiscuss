package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
)

type IQuestionService interface {
	CreateQuestion(ctx context.Context)
}

type QuestionService struct {
	repo repository.IQuestionRepository
}

func NewQuestionService(repo repository.IQuestionRepository) IQuestionService {
	return &QuestionService{repo: repo}
}
