package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (cr *CommentRepository) GetCommentListByAnswerID(ctx context.Context, answerID int) ([]dto.CommentHighlight, error) {
	if answerID == 0 {
		return nil, errors.New("answer id can not be empty")
	}

	Comments := []dto.CommentHighlight{}

	querySelect := `
		%s, 
		%s as author, 
		%s as author_id, 
		%s as comment, 
		%s, 
		%s, 
		%s,
		%s
	`
	querySelect = fmt.Sprintf(
		querySelect,
		model.COMMENT_ID, model.USER_USERNAME, model.USER_ID, model.COMMENT_BODY,
		model.COMMENT_STATISTIC_THUMBS_UP, model.COMMENT_STATISTIC_THUMBS_DOWN,
		model.COMMENT_CREATED_AT, model.COMMENT_UPDATED_AT,
	)

	queryJoinUser := "inner join %s on %s = %s"
	queryJoinUser = fmt.Sprintf(
		queryJoinUser,
		model.USER_TABLE_NAME, model.COMMENT_USER_ID, model.USER_ID,
	)

	queryJoinCommentStatistic := "inner join %s on %s = %s"
	queryJoinCommentStatistic = fmt.Sprintf(
		queryJoinCommentStatistic,
		model.COMMENT_STATISTIC_TABLE_NAME, model.COMMENT_ID,
		model.COMMENT_STATISTIC_COMMENT_ID,
	)

	answerIDEqualTo := fmt.Sprintf("%s = ?", model.COMMENT_ANSWER_ID)

	q := cr.getTrOrDB(ctx).Select(querySelect).
		Table(model.COMMENT_TABLE_NAME).
		Joins(queryJoinUser).
		Joins(queryJoinCommentStatistic).
		Where(answerIDEqualTo, answerID).
		Find(&Comments)

	if q.Error != nil {
		return nil, q.Error
	}

	return Comments, nil
}
