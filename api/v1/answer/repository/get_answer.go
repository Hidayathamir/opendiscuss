package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) GetAnswerByID(ctx context.Context, answerID int) (dto.AnswerHighligh, error) {
	if answerID == 0 {
		return dto.AnswerHighligh{}, errors.New("answer id can not be empty")
	}

	answer := dto.AnswerHighligh{}

	querySelect := `
		%s,
		%s as author,
		%s as author_id,
		%s as answer,
		%s,
		%s,
		%s,
		%s
	`
	querySelect = fmt.Sprintf(
		querySelect,
		model.ANSWER_ID, model.USER_USERNAME, model.USER_ID, model.ANSWER_BODY,
		model.ANSWER_STATISTIC_THUMBS_UP, model.ANSWER_STATISTIC_THUMBS_DOWN,
		model.ANSWER_CREATED_AT, model.ANSWER_UPDATED_AT,
	)

	queryJoinUser := fmt.Sprintf(
		"inner join %s on %s = %s",
		model.USER_TABLE_NAME, model.ANSWER_USER_ID, model.USER_ID,
	)

	queryJoinAnswerStatistic := fmt.Sprintf(
		"inner join %s on %s = %s",
		model.ANSWER_STATISTIC_TABLE_NAME, model.ANSWER_ID,
		model.ANSWER_STATISTIC_ANSWER_ID,
	)

	q := ar.getTrOrDB(ctx).
		Model(&model.Answer{}).
		Select(querySelect).
		Table(model.ANSWER_TABLE_NAME).
		Joins(queryJoinUser).
		Joins(queryJoinAnswerStatistic).
		Where(model.ANSWER_ID, answerID).
		First(&answer)

	if q.Error != nil {
		return dto.AnswerHighligh{}, q.Error
	}

	return answer, nil
}
