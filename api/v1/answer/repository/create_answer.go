package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) CreateAnswer(ctx context.Context, answer model.Answer) (int, error) {
	q := ar.getTrOrDB(ctx).Table(model.ANSWER_TABLE_NAME).Create(&answer)
	if q.Error != nil {
		return 0, q.Error
	}
	return answer.ID, nil
}
