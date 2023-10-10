package model

import (
	"time"

	"github.com/Hidayathamir/opendiscuss/constant"
)

type UserCommentVote struct {
	ID           int                 `gorm:"column:id"`
	UserID       int                 `gorm:"column:user_id"`
	CommentID    int                 `gorm:"column:comment_id"`
	VoteOptionID constant.VoteOption `gorm:"column:vote_option_id"`
	CreatedAt    time.Time           `gorm:"column:created_at"`
	UpdatedAt    time.Time           `gorm:"column:updated_at"`
}

func (UserCommentVote) TableName() string {
	return "user_comment_votes"
}

const (
	USER_COMMENT_VOTES_TABLE_NAME     = "user_comment_votes"
	USER_COMMENT_VOTES_ID             = "user_comment_votes.id"
	USER_COMMENT_VOTES_USER_ID        = "user_comment_votes.user_id"
	USER_COMMENT_VOTES_COMMENT_ID     = "user_comment_votes.comment_id"
	USER_COMMENT_VOTES_VOTE_OPTION_ID = "user_comment_votes.vote_option_id"
	USER_COMMENT_VOTES_CREATED_AT     = "user_comment_votes.created_at"
	USER_COMMENT_VOTES_UPDATED_AT     = "user_comment_votes.updated_at"
)
