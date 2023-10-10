package model

import "time"

type ParentChildComment struct {
	ID              int       `gorm:"column:id"`
	ParentCommentID int       `gorm:"column:parent_comment_id"`
	ChildCommentID  int       `gorm:"column:child_comment_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (ParentChildComment) TableName() string {
	return "parent_child_comments"
}

const (
	PARENT_CHILD_COMMENT_TABLE_NAME        = "parent_child_comments"
	PARENT_CHILD_COMMENT_ID                = "parent_child_comments.id"
	PARENT_CHILD_COMMENT_PARENT_COMMENT_ID = "parent_child_comments.parent_comment_id"
	PARENT_CHILD_COMMENT_CHILD_COMMENT_ID  = "parent_child_comments.child_comment_id"
	PARENT_CHILD_COMMENT_CREATED_AT        = "parent_child_comments.created_at"
	PARENT_CHILD_COMMENT_UPDATED_AT        = "parent_child_comments.updated_at"
)
