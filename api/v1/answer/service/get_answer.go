package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/pkg/errors"
)

func (as *AnswerService) GetAnswerByID(ctx context.Context, answerID int) (dto.AnswerHighligh, error) {
	if answerID == 0 {
		return dto.AnswerHighligh{}, errors.New("answer id can not be empty")
	}

	answer, err := as.repo.GetAnswerByID(ctx, answerID)
	if err != nil {
		return dto.AnswerHighligh{}, errors.Wrap(err, "error get answer by id")
	}

	return answer, nil
}
