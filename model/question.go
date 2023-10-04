package model

import (
	"time"
)

type Question struct {
	ID        int       `gorm:"column:id"`
	UserID    int       `gorm:"column:user_id"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Question) TableName() string {
	return "questions"
}

const (
	QUESTION_TABLE_NAME = "questions"
	QUESTION_ID         = "id"
	QUESTION_USER_ID    = "user_id"
	QUESTION_BODY       = "body"
	QUESTION_CREATED_AT = "created_at"
	QUESTION_UPDATED_AT = "updated_at"
)
