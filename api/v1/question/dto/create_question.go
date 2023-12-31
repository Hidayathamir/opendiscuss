package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateQuestion struct {
	UserID   int    `json:"-"`
	Title    string `json:"title"`
	Question string `json:"question"`
}

func (r ReqCreateQuestion) ToModelQuestion() model.Question {
	return model.Question{
		UserID: r.UserID,
		Title:  r.Title,
		Body:   r.Question,
	}
}

func (r ReqCreateQuestion) Validate() error {
	if r.UserID == 0 {
		return errors.New("user_id can not be empty")
	}
	if r.Title == "" {
		return errors.New("title can not be empty")
	}
	if r.Question == "" {
		return errors.New("question can not be empty")
	}
	return nil
}

type ResCreateQuestion struct {
	QuestionID int `json:"question_id"`
}
