package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) CreateUserAnswerVote(ctx context.Context, userAnswerVote model.UserAnswerVote) (int, error) {
	q := ar.getTrOrDB(ctx).
		Table(model.USER_ANSWER_VOTES_TABLE_NAME).
		Create(&userAnswerVote)

	if q.Error != nil {
		return 0, q.Error
	}

	return userAnswerVote.ID, nil
}
