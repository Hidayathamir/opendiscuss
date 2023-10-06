package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) CreateQuestion(ctx context.Context, question model.Question) (int, error) {
	q := qr.getTrOrDB(ctx).Table(model.QUESTION_TABLE_NAME).Create(&question)
	if q.Error != nil {
		return 0, q.Error
	}
	return question.ID, nil
}
