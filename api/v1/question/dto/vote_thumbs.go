package dto

import (
	"errors"
)

type ReqVoteThumbs struct {
	UserID     int
	QuestionID int
}

func (r ReqVoteThumbs) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	return nil
}

func (r ReqVoteThumbs) ToReqCreateUserQuestionVote() ReqCreateUserQuestionVote {
	return ReqCreateUserQuestionVote{
		UserID:     r.UserID,
		QuestionID: r.QuestionID,
	}
}

type ResVoteThumbs struct {
	QuestionID int `json:"question_id"`
}
