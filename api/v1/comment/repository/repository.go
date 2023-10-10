package repository

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	CreateCommentStatistic(ctx context.Context, commentStatistic model.CommentStatistic) (int, error)
	GetCommentByID(ctx context.Context, ID int) (dto.CommentHighlight, error)
	GetCommentListByAnswerID(ctx context.Context, answerID int) ([]dto.CommentHighlight, error)
	GetUserCommentVoteByUserIDAndCommentID(ctx context.Context, req dto.ReqGetUserCommentVoteByUserIDAndCommentID) (model.UserCommentVote, error)
	CreateUserCommentVote(ctx context.Context, userCommentVote model.UserCommentVote) (int, error)
	UpdateUserCommentVoteColumnVoteOptionIDByID(ctx context.Context, req dto.ReqUpdateUserCommentVoteColumnVoteOptionIDByID) (int, error)
	IncrementCommentStatisticColumnThumbsUpByCommentID(ctx context.Context, commentID int) (int, error)
	IncrementCommentStatisticColumnThumbsDownByCommentID(ctx context.Context, commentID int) (int, error)
	DecrementCommentStatisticColumnThumbsUpByCommentID(ctx context.Context, commentID int) (int, error)
	DecrementCommentStatisticColumnThumbsDownByCommentID(ctx context.Context, commentID int) (int, error)
	UpdateCommentByID(ctx context.Context, req dto.ReqUpdateCommentByID) (int, error)
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
