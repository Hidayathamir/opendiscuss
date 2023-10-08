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

func (AnswerStatistic) TableName() string {
	return "answer_statistics"
}

const (
	ANSWER_STATISTIC_TABLE_NAME  = "answer_statistics"
	ANSWER_STATISTIC_ID          = "answer_statistics.id"
	ANSWER_STATISTIC_ANSWER_ID   = "answer_statistics.answer_id"
	ANSWER_STATISTIC_THUMBS_UP   = "answer_statistics.thumbs_up"
	ANSWER_STATISTIC_THUMBS_DOWN = "answer_statistics.thumbs_down"
	ANSWER_STATISTIC_CREATED_AT  = "answer_statistics.created_at"
	ANSWER_STATISTIC_UPDATED_AT  = "answer_statistics.updated_at"
)
