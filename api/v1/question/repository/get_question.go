package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) GetQuestionByID(ctx context.Context, ID int) (dto.QuestionHighlight, error) {
	if ID == 0 {
		return dto.QuestionHighlight{}, errors.New("id can not be empty")
	}

	question := dto.QuestionHighlight{}

	querySelect := `
		%s, 
		%s as author, 
		%s as author_id, 
		%s as question, 
		%s, 
		%s, 
		%s,
		%s
	`
	querySelect = fmt.Sprintf(
		querySelect,
		model.QUESTION_ID, model.USER_USERNAME, model.USER_ID, model.QUESTION_BODY,
		model.QUESTION_STATISTIC_THUMBS_UP, model.QUESTION_STATISTIC_THUMBS_DOWN,
		model.QUESTION_CREATED_AT, model.QUESTION_UPDATED_AT,
	)

	queryJoinUser := "inner join %s on %s = %s"
	queryJoinUser = fmt.Sprintf(
		queryJoinUser,
		model.USER_TABLE_NAME, model.QUESTION_USER_ID, model.USER_ID,
	)

	queryJoinQuestionStatistic := "inner join %s on %s = %s"
	queryJoinQuestionStatistic = fmt.Sprintf(
		queryJoinQuestionStatistic,
		model.QUESTION_STATISTIC_TABLE_NAME, model.QUESTION_ID,
		model.QUESTION_STATISTIC_QUESTION_ID,
	)

	q := qr.getTrOrDB(ctx).
		Model(&model.Question{}).
		Select(querySelect).
		Table(model.QUESTION_TABLE_NAME).
		Joins(queryJoinUser).
		Joins(queryJoinQuestionStatistic).
		Where(model.QUESTION_ID, ID).
		First(&question)

	if q.Error != nil {
		return dto.QuestionHighlight{}, q.Error
	}

	return question, nil
}
