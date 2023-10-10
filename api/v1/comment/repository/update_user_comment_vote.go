package repository

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
)

func (cr *CommentRepository) UpdateUserCommentVoteColumnVoteOptionIDByID(ctx context.Context, req dto.ReqUpdateUserCommentVoteColumnVoteOptionIDByID) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	idEqualTo := fmt.Sprintf("%s = ?", model.USER_COMMENT_VOTES_COMMENT_ID)

	q := cr.getTrOrDB(ctx).
		Table(model.USER_COMMENT_VOTES_TABLE_NAME).
		Where(idEqualTo, req.ID).
		Update(model.USER_COMMENT_VOTES_VOTE_OPTION_ID, req.VoteOptionID)

	if q.Error != nil {
		return 0, q.Error
	}

	return req.ID, nil
}
