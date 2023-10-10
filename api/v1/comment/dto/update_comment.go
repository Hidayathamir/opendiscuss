package dto

import "errors"

type ReqUpdateCommentByID struct {
	UserID    int    `json:"-"`
	CommentID int    `json:"-"`
	Comment   string `json:"comment"`
}

func (r ReqUpdateCommentByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.CommentID == 0 {
		return errors.New("comment id can not be empty")
	}
	if r.Comment == "" {
		return errors.New("comment can not be empty")
	}
	return nil
}

type ResUpdateCommentByID struct {
	CommentID int `json:"comment_id"`
}
