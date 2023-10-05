package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) CreateQuestionStatistic(ctx context.Context, questionStatistic model.QuestionStatistic) (int, error) {
	q := qr.db.Table(model.QUESTION_STATISTIC_TABLE_NAME).
		Create(&questionStatistic)

	if q.Error != nil {
		return 0, q.Error
	}

	return questionStatistic.ID, nil
}
