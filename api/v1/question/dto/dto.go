package dto

import "time"

type QuestionHighlight struct {
	ID         int       `json:"id"`
	Author     string    `json:"author"`
	AuthorID   int       `json:"author_id"`
	Question   string    `json:"question"`
	ThumbsUp   int       `json:"thumbs_up"`
	ThumbsDown int       `json:"thumbs_down"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
