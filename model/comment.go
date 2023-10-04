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
