package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) GetAnswerListByQuestionID(ctx context.Context, questoinID int) ([]dto.AnswerHighligh, error) {
	if questoinID == 0 {
		return nil, errors.New("question id can not be empty")
	}

	answers := []dto.AnswerHighligh{}

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

	questionIDEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_QUESTION_ID)

	q := ar.getTrOrDB(ctx).
		Select(querySelect).
		Table(model.ANSWER_TABLE_NAME).
		Joins(queryJoinUser).
		Joins(queryJoinAnswerStatistic).
		Where(questionIDEqualTo, questoinID).
		Find(&answers)

	if q.Error != nil {
		return nil, q.Error
	}

	return answers, nil
}
