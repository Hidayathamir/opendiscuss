package service

import (
	"context"
	"strings"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (qs *QuestionService) CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_BODY_VALIDATE)
	}

	var questionID int
	var err error

	err = qs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		questionID, err = qs.repo.CreateQuestion(ctx, req.ToModelQuestion())
		if err != nil {
			err = errors.Wrap(err, "error create question")
			if strings.Contains(err.Error(), "REFERENCES `users` (`id`)") {
				return errors.Wrap(err, "user not found")
			}
			return err
		}

		questionStatistic := model.QuestionStatistic{
			QuestionID: questionID,
		}
		_, err = qs.repo.CreateQuestionStatistic(ctx, questionStatistic)
		if err != nil {
			return errors.Wrap(err, "error create question statistic")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create question and create question statistic")
	}

	return questionID, err
}
