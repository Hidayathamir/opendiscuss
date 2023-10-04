package model

import "time"

type AnswerStatistic struct {
	ID         int       `gorm:"column:id"`
	AnswerID   int       `gorm:"column:answer_id"`
	ThumbsUp   int       `gorm:"column:thumbs_up"`
	ThumbsDown int       `gorm:"column:thumbs_down"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
