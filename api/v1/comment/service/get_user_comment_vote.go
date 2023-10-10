package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (cs *CommentService) getUserCommentVoteByUserIDAndCommentID(ctx context.Context, req dto.ReqGetUserCommentVoteByUserIDAndCommentID) (model.UserCommentVote, error) {
	if err := req.Validate(); err != nil {
		err = errors.Wrap(err, constant.ERR_REQ_VALIDATE)
		return model.UserCommentVote{}, err
	}

	userCommentVote, err := cs.repo.GetUserCommentVoteByUserIDAndCommentID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error get user comment vote by user_id and comment_id")
		return model.UserCommentVote{}, err
	}

	return userCommentVote, nil
}
