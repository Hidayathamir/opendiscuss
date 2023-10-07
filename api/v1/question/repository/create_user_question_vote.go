package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) CreateUserQuestionVote(ctx context.Context, userQuestionVote model.UserQuestionVote) (int, error) {
	q := qr.getTrOrDB(ctx).
		Table(model.USER_QUESTION_VOTES_TABLE_NAME).
		Create(&userQuestionVote)

	if q.Error != nil {
		return 0, q.Error
	}

	return userQuestionVote.ID, nil
}
