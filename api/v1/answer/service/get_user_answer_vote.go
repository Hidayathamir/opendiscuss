package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (as *AnswerService) getUserAnswerVoteByUserIDAndAnswerID(ctx context.Context, req dto.ReqGetUserAnswerVoteByUserIDAndAnswerID) (model.UserAnswerVote, error) {
	if err := req.Validate(); err != nil {
		err = errors.Wrap(err, constant.ERR_REQ_VALIDATE)
		return model.UserAnswerVote{}, err
	}

	userAnswerVote, err := as.repo.GetUserAnswerVoteByUserIDAndAnswerID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error get user answer vote by user_id and answer_id")
		return model.UserAnswerVote{}, err
	}

	return userAnswerVote, nil
}
