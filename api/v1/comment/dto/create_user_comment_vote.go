package dto

import (
	"errors"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
)

type ReqCreateUserCommentVote struct {
	UserID       int
	CommentID    int
	VoteOptionID constant.VoteOption
}

func (r ReqCreateUserCommentVote) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id can not be empty")
	}
	if r.CommentID == 0 {
		return errors.New("comment id can not be empty")
	}
	if r.VoteOptionID == 0 {
		return errors.New("vote option can not be empty")
	}
	return nil
}

func (r ReqCreateUserCommentVote) ToModelUserCommentVote() model.UserCommentVote {
	return model.UserCommentVote{
		UserID:       r.UserID,
		CommentID:    r.CommentID,
		VoteOptionID: r.VoteOptionID,
	}
}
