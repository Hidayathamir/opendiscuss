package dto

import "time"

type QuestionHighlight struct {
	Author     string    `json:"author"`
	Question   string    `json:"question"`
	ThumbsUp   string    `json:"thumbs_up"`
	ThumbsDown string    `json:"thumbs_down"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
