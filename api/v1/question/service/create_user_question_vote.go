package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (qs *QuestionService) createUserQuestionVote(ctx context.Context, req dto.ReqCreateUserQuestionVote) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userQuestionVoteID, err := qs.repo.CreateUserQuestionVote(ctx, req.ToModelUserQuestionVote())
	if err != nil {
		return 0, errors.Wrap(err, "error create user question vote")
	}

	return userQuestionVoteID, nil
}
