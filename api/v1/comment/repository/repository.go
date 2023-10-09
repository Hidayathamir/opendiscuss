package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	CreateCommentStatistic(ctx context.Context, commentStatistic model.CommentStatistic) (int, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db: db}
}

func (qr *CommentRepository) getTrOrDB(ctx context.Context) *gorm.DB {
	isHasTransaction := ctx.Value(utils.CtxKey) != nil
	if isHasTransaction {
		return ctx.Value(utils.CtxKey).(*gorm.DB)
	}
	return qr.db
}
