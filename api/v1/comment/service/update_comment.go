package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (cs *CommentService) UpdateCommentByID(ctx context.Context, req dto.ReqUpdateCommentByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	commentHighlight, err := cs.repo.GetCommentByID(ctx, req.CommentID)
	if err != nil {
		return 0, errors.Wrap(err, "error get comment by id")
	}

	if commentHighlight.AuthorID != req.UserID {
		return 0, errors.New(constant.ERR_UNAUTHORIZED)
	}

	commentID, err := cs.repo.UpdateCommentByID(ctx, req)
	if err != nil {
		return 0, errors.Wrap(err, "error update comment by id")
	}

	return commentID, nil
}
