package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/constant"
)

type ReqUpdateUserQuestionVoteColumnVoteOptionIDByID struct {
	ID           int                 `json:"-"`
	VoteOptionID constant.VoteOption `json:"-"`
}

func (r ReqUpdateUserQuestionVoteColumnVoteOptionIDByID) Validate() error {
	if r.ID == 0 {
		return errors.New("id can not be empty")
	}
	if r.VoteOptionID == 0 {
		return errors.New("vote option id can not be empty")
	}
	return nil
}
