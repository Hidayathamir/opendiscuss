package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
	"gorm.io/gorm"
)

type IQuestionRepository interface {
	CreateQuestion(ctx context.Context, question model.Question) (int, error)
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{db: db}
}
