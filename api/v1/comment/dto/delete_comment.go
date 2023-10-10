package dto

import "errors"

type ReqDeleteCommentByID struct {
	UserID    int
	CommentID int
}

func (r ReqDeleteCommentByID) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.CommentID == 0 {
		return errors.New("comment id can not be empty")
	}
	return nil
}

type ResDeleteCommentByID struct {
	CommentID int `json:"comment_id"`
}
