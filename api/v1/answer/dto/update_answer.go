package dto

import "errors"

type ReqUpdateAnswerByID struct {
	UserID   int    `json:"-"`
	AnswerID int    `json:"-"`
	Answer   string `json:"answer"`
}

func (r ReqUpdateAnswerByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	if r.Answer == "" {
		return errors.New("answer can not be empty")
	}
	return nil
}

type ResUpdateAnswerByID struct {
	AnswerID int `json:"answer_id"`
}
