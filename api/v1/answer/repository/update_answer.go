package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/model"
)

func (ar *AnswerRepository) UpdateAnswerByID(ctx context.Context, req dto.ReqUpdateAnswerByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	idEqualTo := fmt.Sprintf("%s = ?", model.ANSWER_ID)
	q := ar.getTrOrDB(ctx).
		Table(model.ANSWER_TABLE_NAME).
		Where(idEqualTo, req.AnswerID).
		Update(model.ANSWER_BODY, req.Answer)

	if q.Error != nil {
		return 0, q.Error
	}

	return req.AnswerID, nil
}
