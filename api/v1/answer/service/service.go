package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/answer/repository"
	"github.com/Hidayathamir/opendiscuss/utils"
)

type IAnswerService interface {
	CreateAnswer(ctx context.Context, req dto.ReqCreateAnswer) (int, error)
	GetAnswerListByQuestionID(ctx context.Context, questionID int) ([]dto.AnswerHighligh, error)
	GetAnswerByID(ctx context.Context, answerID int) (dto.AnswerHighligh, error)
	VoteThumbsUp(ctx context.Context, req dto.ReqVoteThumbs) (int, error)
	VoteThumbsDown(ctx context.Context, req dto.ReqVoteThumbs) (int, error)
	UpdateAnswerByID(ctx context.Context, req dto.ReqUpdateAnswerByID) (int, error)
	DeleteAnswerByID(ctx context.Context, req dto.ReqDeleteAnswerByID) (int, error)
}

type AnswerService struct {
	repo      repository.IAnswerRepository
	trManager utils.ITransactionManager
}

func NewAnswerService(repo repository.IAnswerRepository, trManager utils.ITransactionManager) IAnswerService {
	return &AnswerService{repo: repo, trManager: trManager}
}
