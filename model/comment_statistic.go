package model

import "time"

type CommentStatistic struct {
	ID         int       `gorm:"column:id"`
	CommentID  int       `gorm:"column:comment_id"`
	ThumbsUp   int       `gorm:"column:thumbs_up"`
	ThumbsDown int       `gorm:"column:thumbs_down"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
