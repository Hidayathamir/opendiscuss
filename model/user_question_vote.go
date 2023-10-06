package model

import (
	"time"

	"github.com/Hidayathamir/opendiscuss/constant"
)

type UserQuestionVote struct {
	ID           int                 `gorm:"column:id"`
	UserID       int                 `gorm:"column:user_id"`
	QuestionID   int                 `gorm:"column:question_id"`
	VoteOptionID constant.VoteOption `gorm:"column:vote_option_id"`
	CreatedAt    time.Time           `gorm:"column:created_at"`
	UpdatedAt    time.Time           `gorm:"column:updated_at"`
}

func (UserQuestionVote) TableName() string {
	return "user_question_votes"
}

const (
	USER_QUESTION_VOTES_TABLE_NAME     = "user_question_votes"
	USER_QUESTION_VOTES_ID             = "user_question_votes.id"
	USER_QUESTION_VOTES_USER_ID        = "user_question_votes.user_id"
	USER_QUESTION_VOTES_QUESTION_ID    = "user_question_votes.question_id"
	USER_QUESTION_VOTES_VOTE_OPTION_ID = "user_question_votes.vote_option_id"
	USER_QUESTION_VOTES_CREATED_AT     = "user_question_votes.created_at"
	USER_QUESTION_VOTES_UPDATED_AT     = "user_question_votes.updated_at"
)
