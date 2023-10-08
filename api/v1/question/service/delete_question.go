package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (qs *QuestionService) DeleteQuestionByID(ctx context.Context, req dto.ReqDeleteQuestionByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	questionHighlight, err := qs.repo.GetQuestionByID(ctx, req.QuestionID)
	if err != nil {
		return 0, errors.Wrap(err, "error get question by id")
	}

	if questionHighlight.AuthorID != req.UserID {
		return 0, errors.New(constant.ERR_UNAUTHORIZED)
	}

	questionID, err := qs.repo.DeleteQuestionByID(ctx, req)
	if err != nil {
		return 0, errors.Wrap(err, "error delete question by id")
	}

	return questionID, nil
}
