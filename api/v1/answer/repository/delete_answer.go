package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) DeleteAnswerByID(ctx context.Context, req dto.ReqDeleteAnswerByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	idEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_ID)
	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_TABLE_NAME).
		Where(idEqualTo, req.AnswerID).
		Delete(&model.Answer{})

	if q.Error != nil {
		return 0, q.Error
	}

	return req.AnswerID, nil
}
