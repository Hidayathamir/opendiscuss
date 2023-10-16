package model

import "time"

type QuestionStatistic struct {
	ID          int       `gorm:"column:id"`
	QuestionID  int       `gorm:"column:question_id"`
	ThumbsUp    int       `gorm:"column:thumbs_up"`
	ThumbsDown  int       `gorm:"column:thumbs_down"`
	AnswerCount int       `gorm:"column:answer_count"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (QuestionStatistic) TableName() string {
	return "question_statistics"
}

const (
	QUESTION_STATISTIC_TABLE_NAME   = "question_statistics"
	QUESTION_STATISTIC_ID           = "question_statistics.id"
	QUESTION_STATISTIC_QUESTION_ID  = "question_statistics.question_id"
	QUESTION_STATISTIC_THUMBS_UP    = "question_statistics.thumbs_up"
	QUESTION_STATISTIC_THUMBS_DOWN  = "question_statistics.thumbs_down"
	QUESTION_STATISTIC_ANSWER_COUNT = "question_statistics.answer_count"
	QUESTION_STATISTIC_CREATED_AT   = "question_statistics.created_at"
	QUESTION_STATISTIC_UPDATED_AT   = "question_statistics.updated_at"
)
