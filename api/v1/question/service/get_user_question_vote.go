package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (qs *QuestionService) getUserQuestionVoteByUserIDAndQuestionID(ctx context.Context, req dto.ReqGetUserQuestionVoteByUserIDAndQuestionID) (model.UserQuestionVote, error) {
	if err := req.Validate(); err != nil {
		err = errors.Wrap(err, constant.ERR_REQ_VALIDATE)
		return model.UserQuestionVote{}, err
	}

	userQuestionVote, err := qs.repo.GetUserQuestionVoteByUserIDAndQuestionID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error get user question vote by user_id and question_id")
		return model.UserQuestionVote{}, err
	}

	return userQuestionVote, nil
}
