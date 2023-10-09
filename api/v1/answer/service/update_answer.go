package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (as *AnswerService) UpdateAnswerByID(ctx context.Context, req dto.ReqUpdateAnswerByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	answerHighlight, err := as.repo.GetAnswerByID(ctx, req.AnswerID)
	if err != nil {
		return 0, errors.Wrap(err, "error get answer by id")
	}

	if answerHighlight.AuthorID != req.UserID {
		return 0, errors.New(constant.ERR_UNAUTHORIZED)
	}

	answerID, err := as.repo.UpdateAnswerByID(ctx, req)
	if err != nil {
		return 0, errors.Wrap(err, "error update answer by id")
	}

	return answerID, nil
}
