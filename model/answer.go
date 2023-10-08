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

func (Answer) TableName() string {
	return "answers"
}

const (
	ANSWER_TABLE_NAME  = "answers"
	ANSWER_ID          = "answers.id"
	ANSWER_QUESTION_ID = "answers.question_id"
	ANSWER_USER_ID     = "answers.user_id"
	ANSWER_BODY        = "answers.body"
	ANSWER_CREATED_AT  = "answers.created_at"
	ANSWER_UPDATED_AT  = "answers.updated_at"
)
