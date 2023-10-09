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

func (UserAnswerVote) TableName() string {
	return "user_answer_votes"
}

const (
	USER_ANSWER_VOTES_TABLE_NAME     = "user_answer_votes"
	USER_ANSWER_VOTES_ID             = "user_answer_votes.id"
	USER_ANSWER_VOTES_USER_ID        = "user_answer_votes.user_id"
	USER_ANSWER_VOTES_ANSWER_ID      = "user_answer_votes.answer_id"
	USER_ANSWER_VOTES_VOTE_OPTION_ID = "user_answer_votes.vote_option_id"
	USER_ANSWER_VOTES_CREATED_AT     = "user_answer_votes.created_at"
	USER_ANSWER_VOTES_UPDATED_AT     = "user_answer_votes.updated_at"
)
