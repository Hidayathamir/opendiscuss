package dto

import "errors"

type ReqGetUserQuestionVoteByUserIDAndQuestionID struct {
	UserID     int `json:"-"`
	QuestionID int `json:"-"`
}

func (r ReqGetUserQuestionVoteByUserIDAndQuestionID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	return nil
}
