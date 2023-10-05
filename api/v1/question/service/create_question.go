package service

import (
	"context"
	"strings"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (qs *QuestionService) CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_BODY_VALIDATE)
	}

	// TODO: CreateQuestion & CreateQuestionStatistic should be in 1 transaction
	questionID, err := qs.repo.CreateQuestion(ctx, req.ToModelQuestion())
	if err != nil {
		err = errors.Wrap(err, "error create question")
		if strings.Contains(err.Error(), "REFERENCES `users` (`id`)") {
			return 0, errors.Wrap(err, "user not found")
		}
		return 0, err
	}

	_, err = qs.repo.CreateQuestionStatistic(ctx, req.ToModelQuestionStatistic(questionID))
	if err != nil {
		return 0, errors.Wrap(err, "error create question statistic")
	}

	return questionID, err
}
