package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (ar *CommentRepository) GetUserCommentVoteByUserIDAndCommentID(ctx context.Context, req dto.ReqGetUserCommentVoteByUserIDAndCommentID) (model.UserCommentVote, error) {
	if err := req.Validate(); err != nil {
		return model.UserCommentVote{}, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	userCommentVote := model.UserCommentVote{}

	userIDEqualTo := fmt.Sprintf("%s = ?", model.USER_COMMENT_VOTES_USER_ID)
	commentIDEqualTo := fmt.Sprintf("%s = ?", model.USER_COMMENT_VOTES_COMMENT_ID)

	q := ar.getTrOrDB(ctx).
		Model(&model.UserCommentVote{}).
		Table(model.USER_COMMENT_VOTES_TABLE_NAME).
		Where(userIDEqualTo, req.UserID).
		Where(commentIDEqualTo, req.CommentID).
		First(&userCommentVote)

	if q.Error != nil {
		return model.UserCommentVote{}, q.Error
	}

	return userCommentVote, nil
}
