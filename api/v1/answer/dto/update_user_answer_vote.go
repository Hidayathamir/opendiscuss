package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/constant"
)

type ReqUpdateUserAnswerVoteColumnVoteOptionIDByID struct {
	ID           int
	VoteOptionID constant.VoteOption
}

func (r ReqUpdateUserAnswerVoteColumnVoteOptionIDByID) Validate() error {
	if r.ID == 0 {
		return errors.New("id can not be empty")
	}
	if r.VoteOptionID == 0 {
		return errors.New("vote option id can not be empty")
	}
	return nil
}
