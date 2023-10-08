package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"gorm.io/gorm"
)

type IAnswerRepository interface {
	CreateAnswer(ctx context.Context, answer model.Answer) (int, error)
	CreateAnswerStatistic(ctx context.Context, answerStatistic model.AnswerStatistic) (int, error)
	GetAnswerListByQuestionID(ctx context.Context, questoinID int) ([]dto.AnswerHighligh, error)
	GetAnswerByID(ctx context.Context, answerID int) (dto.AnswerHighligh, error)
}

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) IAnswerRepository {
	return &AnswerRepository{db: db}
}

func (ar *AnswerRepository) getTrOrDB(ctx context.Context) *gorm.DB {
	isHasTransaction := ctx.Value(utils.CtxKey) != nil
	if isHasTransaction {
		return ctx.Value(utils.CtxKey).(*gorm.DB)
	}
	return ar.db
}
