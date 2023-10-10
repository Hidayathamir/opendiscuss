package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateComment struct {
	UserID          int    `json:"-"`
	AnswerID        int    `json:"-"`
	Comment         string `json:"comment"`
	ParentCommentID int    `json:"parent_comment_id"`
}

func (r ReqCreateComment) ToModelComment() model.Comment {
	return model.Comment{
		UserID:   r.UserID,
		AnswerID: r.AnswerID,
		Body:     r.Comment,
	}
}

func (r ReqCreateComment) Validate() error {
	if r.UserID == 0 {
		return errors.New("user_id can not be empty")
	}
	if r.AnswerID == 0 {
		return errors.New("answer id can not be empty")
	}
	if r.Comment == "" {
		return errors.New("comment can not be empty")
	}
	return nil
}

type ResCreateComment struct {
	CommentID int `json:"comment_id"`
}
