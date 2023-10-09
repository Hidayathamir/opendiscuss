package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) CreateCommentStatistic(ctx context.Context, commentStatistic model.CommentStatistic) (int, error) {
	q := cr.getTrOrDB(ctx).Table(model.COMMENT_STATISTIC_TABLE_NAME).
		Create(&commentStatistic)

	if q.Error != nil {
		return 0, q.Error
	}

	return commentStatistic.ID, nil
}
