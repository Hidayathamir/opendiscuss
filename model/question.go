package model

import (
	"time"
)

type Question struct {
	ID        int       `gorm:"column:id"`
	UserID    int       `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Question) TableName() string {
	return "questions"
}

const (
	QUESTION_TABLE_NAME = "questions"
	QUESTION_ID         = "questions.id"
	QUESTION_USER_ID    = "questions.user_id"
	QUESTION_TITLE      = "questions.title"
	QUESTION_BODY       = "questions.body"
	QUESTION_CREATED_AT = "questions.created_at"
	QUESTION_UPDATED_AT = "questions.updated_at"
)
