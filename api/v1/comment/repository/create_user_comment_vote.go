package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) CreateUserCommentVote(ctx context.Context, userCommentVote model.UserCommentVote) (int, error) {
	q := cr.getTrOrDB(ctx).
		Table(model.USER_COMMENT_VOTES_TABLE_NAME).
		Create(&userCommentVote)

	if q.Error != nil {
		return 0, q.Error
	}

	return userCommentVote.ID, nil
}
