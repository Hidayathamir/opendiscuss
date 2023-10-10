package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (cs *CommentService) VoteThumbsUp(ctx context.Context, req dto.ReqVoteThumbs) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	return cs.voteThumbs(ctx, req, constant.VoteOptionThumbsUp)
}
