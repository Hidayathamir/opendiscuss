package dto

import "errors"

type ReqDeleteQuestionByID struct {
	UserID     int
	QuestionID int
}

func (r ReqDeleteQuestionByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	return nil
}

type ResDeleteQuestionByID struct {
	QuestionID int `json:"question_id"`
}
