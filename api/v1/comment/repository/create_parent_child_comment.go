package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) CreateParentChildComment(ctx context.Context, parentChildComment model.ParentChildComment) (int, error) {
	q := cr.getTrOrDB(ctx).
		Table(model.PARENT_CHILD_COMMENT_TABLE_NAME).
		Create(&parentChildComment)

	if q.Error != nil {
		return 0, q.Error
	}

	return parentChildComment.ID, nil
}
