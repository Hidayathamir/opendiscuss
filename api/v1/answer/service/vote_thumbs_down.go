package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/pkg/errors"
)

func (as *AnswerService) VoteThumbsDown(ctx context.Context, req dto.ReqVoteThumbs) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	return as.voteThumbs(ctx, req, constant.VoteOptionThumbsDown)
}
