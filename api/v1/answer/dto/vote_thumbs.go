package dto

import "errors"

type ReqVoteThumbs struct {
	UserID   int
	AnswerID int
}

func (r ReqVoteThumbs) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	return nil
}

func (r ReqVoteThumbs) ToReqCreateUserAnswerVote() ReqCreateUserAnswerVote {
	return ReqCreateUserAnswerVote{
		UserID:   r.UserID,
		AnswerID: r.AnswerID,
	}
}

type ResVoteThumbs struct {
	AnswerID int `json:"answer_id"`
}
