package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/pkg/errors"
)

func (cs *CommentService) GetCommentListByAnswerID(ctx context.Context, answerID int) ([]dto.CommentHighlight, error) {
	if answerID == 0 {
		return nil, errors.New("answer id can not be empty")
	}

	comments, err := cs.repo.GetCommentListByAnswerID(ctx, answerID)
	if err != nil {
		return nil, errors.Wrap(err, "error get comment list by answer id")
	}
	return comments, nil
}
