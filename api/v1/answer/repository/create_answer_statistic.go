package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) CreateAnswerStatistic(ctx context.Context, answerStatistic model.AnswerStatistic) (int, error) {
	q := ar.getTrOrDB(ctx).Table(model.ANSWER_STATISTIC_TABLE_NAME).
		Create(&answerStatistic)

	if q.Error != nil {
		return 0, q.Error
	}

	return answerStatistic.ID, nil
}
