package dto

import "errors"

type ReqGetUserAnswerVoteByUserIDAndAnswerID struct {
	UserID   int
	AnswerID int
}

func (r ReqGetUserAnswerVoteByUserIDAndAnswerID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	return nil
}
