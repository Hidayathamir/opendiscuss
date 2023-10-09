package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (qs *AnswerService) DeleteAnswerByID(ctx context.Context, req dto.ReqDeleteAnswerByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	answerHighlight, err := qs.repo.GetAnswerByID(ctx, req.AnswerID)
	if err != nil {
		return 0, errors.Wrap(err, "error get answer by id")
	}

	if answerHighlight.AuthorID != req.UserID {
		return 0, errors.New(constant.ERR_UNAUTHORIZED)
	}

	answerID, err := qs.repo.DeleteAnswerByID(ctx, req)
	if err != nil {
		return 0, errors.Wrap(err, "error delete answer by id")
	}

	return answerID, nil
}
