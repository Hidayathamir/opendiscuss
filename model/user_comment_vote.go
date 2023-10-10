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
