package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) CreateComment(ctx context.Context, comment model.Comment) (int, error) {
	q := cr.getTrOrDB(ctx).Table(model.COMMENT_TABLE_NAME).Create(&comment)
	if q.Error != nil {
		return 0, q.Error
	}
	return comment.ID, nil
}
