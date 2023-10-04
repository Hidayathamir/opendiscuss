package service

import (
	"context"
	"strings"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
)

func (qs *QuestionService) CreateQuestion(ctx context.Context, req dto.ReqCreateQuestion) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	questionID, err := qs.repo.CreateQuestion(ctx, req.ToModelQuestion())
	if err != nil {
		if strings.Contains(err.Error(), "questions_ibfk_1") {
			return 0, constant.USER_NOT_FOUND
		}
		return 0, err
	}

	return questionID, err
}
