package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (ar *AnswerRepository) GetUserAnswerVoteByUserIDAndAnswerID(ctx context.Context, req dto.ReqGetUserAnswerVoteByUserIDAndAnswerID) (model.UserAnswerVote, error) {
	if err := req.Validate(); err != nil {
		return model.UserAnswerVote{}, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userAnswerVote := model.UserAnswerVote{}

	userIDEqualTo := fmt.Sprintf("%s = ?", model.USER_ANSWER_VOTES_USER_ID)
	answerIDEqualTo := fmt.Sprintf("%s = ?", model.USER_ANSWER_VOTES_ANSWER_ID)

	q := ar.getTrOrDB(ctx).
		Model(&model.UserAnswerVote{}).
		Table(model.USER_ANSWER_VOTES_TABLE_NAME).
		Where(userIDEqualTo, req.UserID).
		Where(answerIDEqualTo, req.AnswerID).
		First(&userAnswerVote)

	if q.Error != nil {
		return model.UserAnswerVote{}, q.Error
	}

	return userAnswerVote, nil
}
