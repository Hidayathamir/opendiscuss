package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/model"
	"gorm.io/gorm"
)

func (cr *CommentRepository) IncrementCommentStatisticColumnThumbsUpByCommentID(ctx context.Context, commentID int) (int, error) {
	if commentID == 0 {
		return 0, errors.New("comment id can not be empty")
	}

	commentIDEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_STATISTIC_COMMENT_ID)
	incrementThumbsUp := gorm.Expr(fmt.Sprintf("%s + 1", model.COMMENT_STATISTIC_THUMBS_UP))

	q := cr.getTrOrDB(ctx).
		Table(model.COMMENT_STATISTIC_TABLE_NAME).
		Where(commentIDEqualTo, commentID).
		Update(model.COMMENT_STATISTIC_THUMBS_UP, incrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return commentID, nil
}

func (ar *CommentRepository) IncrementCommentStatisticColumnThumbsDownByCommentID(ctx context.Context, commentID int) (int, error) {
	if commentID == 0 {
		return 0, errors.New("comment id can not be empty")
	}

	commentIDEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_STATISTIC_COMMENT_ID)
	incrementThumbsDown := gorm.Expr(fmt.Sprintf("%s + 1", model.COMMENT_STATISTIC_THUMBS_DOWN))

	q := ar.getTrOrDB(ctx).
		Table(model.COMMENT_STATISTIC_TABLE_NAME).
		Where(commentIDEqualTo, commentID).
		Update(model.COMMENT_STATISTIC_THUMBS_DOWN, incrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return commentID, nil
}

func (ar *CommentRepository) DecrementCommentStatisticColumnThumbsUpByCommentID(ctx context.Context, commentID int) (int, error) {
	if commentID == 0 {
		return 0, errors.New("comment id can not be empty")
	}

	commentIDEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_STATISTIC_COMMENT_ID)
	decrementThumbsUp := gorm.Expr(fmt.Sprintf("%s - 1", model.COMMENT_STATISTIC_THUMBS_UP))

	q := ar.getTrOrDB(ctx).
		Table(model.COMMENT_STATISTIC_TABLE_NAME).
		Where(commentIDEqualTo, commentID).
		Update(model.COMMENT_STATISTIC_THUMBS_UP, decrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return commentID, nil
}

func (ar *CommentRepository) DecrementCommentStatisticColumnThumbsDownByCommentID(ctx context.Context, commentID int) (int, error) {
	if commentID == 0 {
		return 0, errors.New("comment id can not be empty")
	}

	commentIDEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_STATISTIC_COMMENT_ID)
	decrementThumbsDown := gorm.Expr(fmt.Sprintf("%s - 1", model.COMMENT_STATISTIC_THUMBS_DOWN))

	q := ar.getTrOrDB(ctx).
		Table(model.COMMENT_STATISTIC_TABLE_NAME).
		Where(commentIDEqualTo, commentID).
		Update(model.COMMENT_STATISTIC_THUMBS_DOWN, decrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return commentID, nil
}
