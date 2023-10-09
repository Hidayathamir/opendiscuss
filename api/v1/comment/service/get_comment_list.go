package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/pkg/errors"
)

func (cs *CommentService) GetCommentList(ctx context.Context) ([]dto.CommentHighlight, error) {
	comments, err := cs.repo.GetCommentList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error get comment list")
	}
	return comments, nil
}
