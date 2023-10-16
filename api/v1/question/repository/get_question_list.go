package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error) {
	questions := []dto.QuestionHighlight{}

	querySelect := `
		%s, 
		%s as author, 
		%s as author_id, 
		%s, 
		%s as question, 
		%s, 
		%s, 
		%s,
		%s,
		%s
	`
	querySelect = fmt.Sprintf(
		querySelect,
		model.QUESTION_ID, model.USER_USERNAME, model.USER_ID,
		model.QUESTION_TITLE, model.QUESTION_BODY,
		model.QUESTION_STATISTIC_THUMBS_UP, model.QUESTION_STATISTIC_THUMBS_DOWN,
		model.QUESTION_STATISTIC_ANSWER_COUNT,
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

	q := qr.getTrOrDB(ctx).Select(querySelect).
		Table(model.QUESTION_TABLE_NAME).
		Joins(queryJoinUser).
		Joins(queryJoinQuestionStatistic).
		Order("created_at desc").
		Find(&questions)

	if q.Error != nil {
		return nil, q.Error
	}

	return questions, nil
}
