package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateAnswer struct {
	UserID     int    `json:"-"`
	QuestionID int    `json:"-"`
	Answer     string `json:"answer"`
}

func (r ReqCreateAnswer) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.QuestionID == 0 {
		return errors.New("question id can not be empty")
	}
	if r.Answer == "" {
		return errors.New("answer can not be empty")
	}
	return nil
}

func (r ReqCreateAnswer) ToModelAnswer() model.Answer {
	return model.Answer{
		UserID:     r.UserID,
		QuestionID: r.QuestionID,
		Body:       r.Answer,
	}
}

type ResCreateAnswer struct {
	AnswerID int `json:"answer_id"`
}
