package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (qr *QuestionRepository) IncrementQuestionStatisticColumnThumbsUpByQuestionID(ctx context.Context, questionID int) (int, error) {
	if questionID == 0 {
		return 0, errors.New("question id can not be empty")
	}

	questionIDEqualTo := fmt.Sprintf("%s = ?", model.QUESTION_STATISTIC_QUESTION_ID)
	incrementThumbsUp := gorm.Expr(fmt.Sprintf("%s + 1", model.QUESTION_STATISTIC_THUMBS_UP))

	q := qr.getTrOrDB(ctx).
		Table(model.QUESTION_STATISTIC_TABLE_NAME).
		Where(questionIDEqualTo, questionID).
		Update(model.QUESTION_STATISTIC_THUMBS_UP, incrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return questionID, nil
}

func (qr *QuestionRepository) IncrementQuestionStatisticColumnThumbsDownByQuestionID(ctx context.Context, questionID int) (int, error) {
	if questionID == 0 {
		return 0, errors.New("question id can not be empty")
	}

	questionIDEqualTo := fmt.Sprintf("%s = ?", model.QUESTION_STATISTIC_QUESTION_ID)
	incrementThumbsDown := gorm.Expr(fmt.Sprintf("%s + 1", model.QUESTION_STATISTIC_THUMBS_DOWN))

	q := qr.getTrOrDB(ctx).
		Table(model.QUESTION_STATISTIC_TABLE_NAME).
		Where(questionIDEqualTo, questionID).
		Update(model.QUESTION_STATISTIC_THUMBS_DOWN, incrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return questionID, nil
}

func (qr *QuestionRepository) DecrementQuestionStatisticColumnThumbsUpByQuestionID(ctx context.Context, questionID int) (int, error) {
	if questionID == 0 {
		return 0, errors.New("question id can not be empty")
	}

	questionIDEqualTo := fmt.Sprintf("%s = ?", model.QUESTION_STATISTIC_QUESTION_ID)
	decrementThumbsUp := gorm.Expr(fmt.Sprintf("%s - 1", model.QUESTION_STATISTIC_THUMBS_UP))

	q := qr.getTrOrDB(ctx).
		Table(model.QUESTION_STATISTIC_TABLE_NAME).
		Where(questionIDEqualTo, questionID).
		Update(model.QUESTION_STATISTIC_THUMBS_UP, decrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return questionID, nil
}

func (qr *QuestionRepository) DecrementQuestionStatisticColumnThumbsDownByQuestionID(ctx context.Context, questionID int) (int, error) {
	if questionID == 0 {
		return 0, errors.New("question id can not be empty")
	}

	questionIDEqualTo := fmt.Sprintf("%s = ?", model.QUESTION_STATISTIC_QUESTION_ID)
	decrementThumbsDown := gorm.Expr(fmt.Sprintf("%s - 1", model.QUESTION_STATISTIC_THUMBS_DOWN))

	q := qr.getTrOrDB(ctx).
		Table(model.QUESTION_STATISTIC_TABLE_NAME).
		Where(questionIDEqualTo, questionID).
		Update(model.QUESTION_STATISTIC_THUMBS_DOWN, decrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return questionID, nil
}
