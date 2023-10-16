package model

import "time"

type CommentStatistic struct {
	ID           int       `gorm:"column:id"`
	CommentID    int       `gorm:"column:comment_id"`
	ThumbsUp     int       `gorm:"column:thumbs_up"`
	ThumbsDown   int       `gorm:"column:thumbs_down"`
	CommentCount int       `gorm:"column:comment_count"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (CommentStatistic) TableName() string {
	return "comment_statistics"
}

const (
	COMMENT_STATISTIC_TABLE_NAME    = "comment_statistics"
	COMMENT_STATISTIC_ID            = "comment_statistics.id"
	COMMENT_STATISTIC_COMMENT_ID    = "comment_statistics.comment_id"
	COMMENT_STATISTIC_THUMBS_UP     = "comment_statistics.thumbs_up"
	COMMENT_STATISTIC_THUMBS_DOWN   = "comment_statistics.thumbs_down"
	COMMENT_STATISTIC_COMMENT_COUNT = "comment_statistics.comment_count"
	COMMENT_STATISTIC_CREATED_AT    = "comment_statistics.created_at"
	COMMENT_STATISTIC_UPDATED_AT    = "comment_statistics.updated_at"
)
