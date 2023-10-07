package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"gorm.io/gorm"
)

type IQuestionRepository interface {
	CreateQuestion(ctx context.Context, question model.Question) (int, error)
	CreateQuestionStatistic(ctx context.Context, question model.QuestionStatistic) (int, error)
	GetQuestionByID(ctx context.Context, ID int) (dto.QuestionHighlight, error)
	GetQuestionList(ctx context.Context) ([]dto.QuestionHighlight, error)
	CreateUserQuestionVote(ctx context.Context, userQuestionVote model.UserQuestionVote) (int, error)
	GetUserQuestionVoteByUserIDAndQuestionID(ctx context.Context, req dto.ReqGetUserQuestionVoteByUserIDAndQuestionID) (model.UserQuestionVote, error)
	UpdateUserQuestionVoteColumnVoteOptionIDByID(ctx context.Context, req dto.ReqUpdateUserQuestionVoteColumnVoteOptionIDByID) (int, error)
	IncrementQuestionStatisticColumnThumbsUpByQuestionID(ctx context.Context, questionID int) (int, error)
	IncrementQuestionStatisticColumnThumbsDownByQuestionID(ctx context.Context, questionID int) (int, error)
	DecrementQuestionStatisticColumnThumbsUpByQuestionID(ctx context.Context, questionID int) (int, error)
	DecrementQuestionStatisticColumnThumbsDownByQuestionID(ctx context.Context, questionID int) (int, error)
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{db: db}
}

func (qr *QuestionRepository) getTrOrDB(ctx context.Context) *gorm.DB {
	isHasTransaction := ctx.Value(utils.CtxKey) != nil
	if isHasTransaction {
		return ctx.Value(utils.CtxKey).(*gorm.DB)
	}
	return qr.db
}
