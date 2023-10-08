package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/pkg/errors"
)

func (as *AnswerService) GetAnswerListByQuestionID(ctx context.Context, questionID int) ([]dto.AnswerHighligh, error) {
	if questionID == 0 {
		return nil, errors.New("question id can not be empty")
	}

	answers, err := as.repo.GetAnswerListByQuestionID(ctx, questionID)
	if err != nil {
		return nil, errors.New("error get answer list by question id")
	}

	return answers, nil
}
