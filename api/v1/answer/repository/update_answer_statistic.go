package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/model"
	"gorm.io/gorm"
)

func (ar *AnswerRepository) IncrementAnswerStatisticColumnThumbsUpByAnswerID(ctx context.Context, answerID int) (int, error) {
	if answerID == 0 {
		return 0, errors.New("answer id can not be empty")
	}

	answerIDEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_STATISTIC_ANSWER_ID)
	incrementThumbsUp := gorm.Expr(fmt.Sprintf("%s + 1", model.ANSWER_STATISTIC_THUMBS_UP))

	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_STATISTIC_TABLE_NAME).
		Where(answerIDEqualTo, answerID).
		Update(model.ANSWER_STATISTIC_THUMBS_UP, incrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return answerID, nil
}

func (ar *AnswerRepository) IncrementAnswerStatisticColumnThumbsDownByAnswerID(ctx context.Context, answerID int) (int, error) {
	if answerID == 0 {
		return 0, errors.New("answer id can not be empty")
	}

	answerIDEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_STATISTIC_ANSWER_ID)
	incrementThumbsDown := gorm.Expr(fmt.Sprintf("%s + 1", model.ANSWER_STATISTIC_THUMBS_DOWN))

	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_STATISTIC_TABLE_NAME).
		Where(answerIDEqualTo, answerID).
		Update(model.ANSWER_STATISTIC_THUMBS_DOWN, incrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return answerID, nil
}
func (ar *AnswerRepository) DecrementAnswerStatisticColumnThumbsUpByAnswerID(ctx context.Context, answerID int) (int, error) {
	if answerID == 0 {
		return 0, errors.New("answer id can not be empty")
	}

	answerIDEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_STATISTIC_ANSWER_ID)
	decrementThumbsUp := gorm.Expr(fmt.Sprintf("%s - 1", model.ANSWER_STATISTIC_THUMBS_UP))

	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_STATISTIC_TABLE_NAME).
		Where(answerIDEqualTo, answerID).
		Update(model.ANSWER_STATISTIC_THUMBS_UP, decrementThumbsUp)

	if q.Error != nil {
		return 0, q.Error
	}

	return answerID, nil
}

func (ar *AnswerRepository) DecrementAnswerStatisticColumnThumbsDownByAnswerID(ctx context.Context, answerID int) (int, error) {
	if answerID == 0 {
		return 0, errors.New("answer id can not be empty")
	}

	answerIDEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_STATISTIC_ANSWER_ID)
	decrementThumbsDown := gorm.Expr(fmt.Sprintf("%s - 1", model.ANSWER_STATISTIC_THUMBS_DOWN))

	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_STATISTIC_TABLE_NAME).
		Where(answerIDEqualTo, answerID).
		Update(model.ANSWER_STATISTIC_THUMBS_DOWN, decrementThumbsDown)

	if q.Error != nil {
		return 0, q.Error
	}

	return answerID, nil
}
