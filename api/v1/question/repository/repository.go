package repository

import (
	"context"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	CreateQuestion(ctx context.Context)
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{db: db}
}
