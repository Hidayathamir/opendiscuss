package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (as *AnswerService) CreateAnswer(ctx context.Context, req dto.ReqCreateAnswer) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	var answerID int
	var err error

	err = as.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		answerID, err = as.repo.CreateAnswer(ctx, req.ToModelAnswer())
		if err != nil {
			return errors.Wrap(err, "error create answer")
		}

		answerStatistic := model.AnswerStatistic{
			AnswerID: answerID,
		}
		_, err := as.repo.CreateAnswerStatistic(ctx, answerStatistic)
		if err != nil {
			return errors.Wrap(err, "error create answer statistic")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create answer and create answer statistic")
	}

	return answerID, nil
}
