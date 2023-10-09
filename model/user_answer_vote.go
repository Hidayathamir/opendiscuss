package model

import (
	"time"

	"github.com/Hidayathamir/opendiscuss/constant"
)

type UserAnswerVote struct {
	ID           int                 `gorm:"column:id"`
	UserID       int                 `gorm:"column:user_id"`
	AnswerID     int                 `gorm:"column:answer_id"`
	VoteOptionID constant.VoteOption `gorm:"column:vote_option_id"`
	CreatedAt    time.Time           `gorm:"column:created_at"`
	UpdatedAt    time.Time           `gorm:"column:updated_at"`
}
