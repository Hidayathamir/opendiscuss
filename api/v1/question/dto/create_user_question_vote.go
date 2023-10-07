package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateUserQuestionVote struct {
	UserID       int
	QuestionID   int
	VoteOptionID constant.VoteOption
}

func (r ReqCreateUserQuestionVote) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	if r.VoteOptionID == 0 {
		return errors.New("vote option can not be empty")
	}
	return nil
}

func (r ReqCreateUserQuestionVote) ToModelUserQuestionVote() model.UserQuestionVote {
	return model.UserQuestionVote{
		UserID:       r.UserID,
		QuestionID:   r.QuestionID,
		VoteOptionID: r.VoteOptionID,
	}
}
