package model

import "time"

type UserQuestionVote struct {
	ID           int       `gorm:"column:id"`
	UserID       int       `gorm:"column:user_id"`
	QuestionID   int       `gorm:"column:question_id"`
	VoteOptionID int       `gorm:"column:vote_option_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
