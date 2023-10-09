package dto

import "errors"

type ReqDeleteAnswerByID struct {
	UserID   int
	AnswerID int
}

func (r ReqDeleteAnswerByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	return nil
}

type ResDeleteAnswerByID struct {
	AnswerID int `json:"answer_id"`
}
