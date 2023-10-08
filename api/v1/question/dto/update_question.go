package dto

import "errors"

type ReqUpdateQuestionByID struct {
	UserID     int    `json:"-"`
	QuestionID int    `json:"-"`
	Question   string `json:"question"`
}

func (r ReqUpdateQuestionByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	if r.Question == "" {
		return errors.New("question can not be empty")
	}
	return nil
}

type ResUpdateQuestionByID struct {
	QuestionID int `json:"question_id"`
}
