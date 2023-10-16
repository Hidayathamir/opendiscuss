package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (qr *QuestionRepository) UpdateQuestionByID(ctx context.Context, req dto.ReqUpdateQuestionByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	idEqualTo := fmt.Sprintf("%s = ?", model.QUESTION_ID)
	q := qr.getTrOrDB(ctx).
		Table(model.QUESTION_TABLE_NAME).
		Where(idEqualTo, req.QuestionID).
		Updates(model.Question{
			Title: req.Title,
			Body:  req.Question,
		})

	if q.Error != nil {
		return 0, q.Error
	}

	return req.QuestionID, nil
}
