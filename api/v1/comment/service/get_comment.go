package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/pkg/errors"
)

func (cs *CommentService) GetCommentByID(ctx context.Context, ID int) (dto.CommentHighlight, error) {
	if ID == 0 {
		return dto.CommentHighlight{}, errors.New("id can not be empty")
	}

	comment, err := cs.repo.GetCommentByID(ctx, ID)
	if err != nil {
		return dto.CommentHighlight{}, errors.Wrap(err, "error get comment")
	}
	return comment, nil
}
