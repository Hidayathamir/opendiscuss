package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) DeleteCommentByID(ctx context.Context, req dto.ReqDeleteCommentByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	idEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_ID)
	q := cr.getTrOrDB(ctx).
		Table(model.COMMENT_TABLE_NAME).
		Where(idEqualTo, req.CommentID).
		Delete(&model.Comment{})

	if q.Error != nil {
		return 0, q.Error
	}

	return req.CommentID, nil
}
