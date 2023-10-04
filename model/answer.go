package model

import "time"

type Answer struct {
	ID         int       `gorm:"column:id"`
	QuestionID int       `gorm:"column:question_id"`
	UserID     int       `gorm:"column:user_id"`
	Body       string    `gorm:"column:body"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
