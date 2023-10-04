package model

import "time"

type ParentChildComment struct {
	ID              int       `gorm:"column:id"`
	ParentCommentID int       `gorm:"column:parent_comment_id"`
	ChildCommentID  int       `gorm:"column:child_comment_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
