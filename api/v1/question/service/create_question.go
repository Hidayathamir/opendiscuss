package service

import "context"

func (qs *QuestionService) CreateQuestion(ctx context.Context) {
	qs.repo.CreateQuestion(ctx)
}
