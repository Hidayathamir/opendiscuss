package model

import "time"

type UserCommentVote struct {
	ID           int       `gorm:"column:id"`
	UserID       int       `gorm:"column:user_id"`
	CommentID    int       `gorm:"column:comment_id"`
	VoteOptionID int       `gorm:"column:vote_option_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
