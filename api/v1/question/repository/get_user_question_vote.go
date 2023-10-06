package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (qr *QuestionRepository) GetUserQuestionVoteByUserIDAndQuestionID(ctx context.Context, req dto.ReqGetUserQuestionVoteByUserIDAndQuestionID) (model.UserQuestionVote, error) {
	if err := req.Validate(); err != nil {
		return model.UserQuestionVote{}, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userQuestionVote := model.UserQuestionVote{}

	userIDEqualTo := fmt.Sprintf("%s = ?", model.USER_QUESTION_VOTES_USER_ID)
	questionIDEqualTo := fmt.Sprintf("%s = ?", model.USER_QUESTION_VOTES_QUESTION_ID)

	q := qr.getTrOrDB(ctx).
		Table(model.USER_QUESTION_VOTES_TABLE_NAME).
		Where(userIDEqualTo, req.UserID).
		Where(questionIDEqualTo, req.QuestionID).
		First(&userQuestionVote)

	if q.Error != nil {
		return model.UserQuestionVote{}, q.Error
	}

	return userQuestionVote, nil
}
