package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (as *AnswerService) createUserAnswerVote(ctx context.Context, req dto.ReqCreateUserAnswerVote) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userAnswerVoteID, err := as.repo.CreateUserAnswerVote(ctx, req.ToModelUserAnswerVote())
	if err != nil {
		return 0, errors.Wrap(err, "error create user answer vote")
	}

	return userAnswerVoteID, nil
}
