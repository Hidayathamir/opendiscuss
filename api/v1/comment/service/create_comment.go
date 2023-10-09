package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (cs *CommentService) CreateComment(ctx context.Context, req dto.ReqCreateComment) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	var commentID int
	var err error

	err = cs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		commentID, err = cs.repo.CreateComment(ctx, req.ToModelComment())
		if err != nil {
			return errors.Wrap(err, "error create comment")
		}

		CommentStatistic := model.CommentStatistic{
			CommentID: commentID,
		}
		_, err = cs.repo.CreateCommentStatistic(ctx, CommentStatistic)
		if err != nil {
			return errors.Wrap(err, "error create comment statistic")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create comment and create comment statistic")
	}

	return commentID, nil
}
