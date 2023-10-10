package model

import "time"

type Comment struct {
	ID        int       `gorm:"column:id"`
	AnswerID  int       `gorm:"column:answer_id"`
	UserID    int       `gorm:"column:user_id"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Comment) TableName() string {
	return "comments"
}

const (
	COMMENT_TABLE_NAME = "comments"
	COMMENT_ID         = "comments.id"
	COMMENT_ANSWER_ID  = "comments.answer_id"
	COMMENT_USER_ID    = "comments.user_id"
	COMMENT_BODY       = "comments.body"
	COMMENT_CREATED_AT = "comments.created_at"
	COMMENT_UPDATED_AT = "comments.updated_at"
)
