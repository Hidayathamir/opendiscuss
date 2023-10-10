package dto

import "errors"

type ReqVoteThumbs struct {
	UserID    int
	CommentID int
}

func (r ReqVoteThumbs) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.CommentID == 0 {
		return errors.New("comment id can not be empty")
	}
	return nil
}

func (r ReqVoteThumbs) ToReqCreateUserCommentVote() ReqCreateUserCommentVote {
	return ReqCreateUserCommentVote{
		UserID:    r.UserID,
		CommentID: r.CommentID,
	}
}

type ResVoteThumbs struct {
	CommentID int `json:"comment_id"`
}
