package model

import "time"

type UserAnswerVote struct {
	ID         int       `gorm:"column:id"`
	QuestionID int       `gorm:"column:question_id"`
	ThumbsUp   int       `gorm:"column:thumbs_up"`
	ThumbsDown int       `gorm:"column:thumbs_down"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
