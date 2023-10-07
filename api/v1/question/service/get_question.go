package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/pkg/errors"
)

func (qs *QuestionService) GetQuestionByID(ctx context.Context, ID int) (dto.QuestionHighlight, error) {
	if ID == 0 {
		return dto.QuestionHighlight{}, errors.New("id can not be empty")
	}

	question, err := qs.repo.GetQuestionByID(ctx, ID)
	if err != nil {
		return dto.QuestionHighlight{}, errors.Wrap(err, "error get question")
	}
	return question, nil
}
