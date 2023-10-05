package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/model"
	"gorm.io/gorm"
)

type IQuestionRepository interface {
	CreateQuestion(ctx context.Context, question model.Question) (int, error)
	CreateQuestionStatistic(ctx context.Context, question model.QuestionStatistic) (int, error)
	GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error)
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{db: db}
}
