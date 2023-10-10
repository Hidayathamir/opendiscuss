package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/pkg/errors"
)

func (qs *QuestionService) GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error) {
	questions, err := qs.repo.GetQuestionList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error get question list")
	}

	for i := range questions {
		questions[i].ThumbsRate = questions[i].ThumbsUp - questions[i].ThumbsDown
	}

	return questions, nil
}
