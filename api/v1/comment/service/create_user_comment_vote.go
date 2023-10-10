package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (cs *CommentService) createUserCommentVote(ctx context.Context, req dto.ReqCreateUserCommentVote) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userCommentVoteID, err := cs.repo.CreateUserCommentVote(ctx, req.ToModelUserCommentVote())
	if err != nil {
		return 0, errors.Wrap(err, "error create user comment vote")
	}

	return userCommentVoteID, nil
}
