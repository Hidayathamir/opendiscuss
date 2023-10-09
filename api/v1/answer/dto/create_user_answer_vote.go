package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateUserAnswerVote struct {
	UserID       int
	AnswerID     int
	VoteOptionID constant.VoteOption
}

func (r ReqCreateUserAnswerVote) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	if r.VoteOptionID == 0 {
		return errors.New("vote option can not be empty")
	}
	return nil
}

func (r ReqCreateUserAnswerVote) ToModelUserAnswerVote() model.UserAnswerVote {
	return model.UserAnswerVote{
		UserID:       r.UserID,
		AnswerID:     r.AnswerID,
		VoteOptionID: r.VoteOptionID,
	}
}
