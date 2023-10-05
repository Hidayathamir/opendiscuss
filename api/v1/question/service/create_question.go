package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (qs *QuestionService) CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_BODY_VALIDATE)
	}

	questionID, err := qs.repo.CreateQuestion(ctx, req.ToModelQuestion())
	if err != nil {
		return 0, errors.Wrap(err, "error create question")
	}

	return questionID, err
}
