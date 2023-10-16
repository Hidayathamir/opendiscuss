package dto

import "time"

type AnswerHighligh struct {
	ID           int       `json:"id"`
	Author       string    `json:"author"`
	AuthorID     int       `json:"author_id"`
	Answer       string    `json:"answer"`
	ThumbsRate   int       `json:"thumbs_rate"`
	ThumbsUp     int       `json:"thumbs_up"`
	ThumbsDown   int       `json:"thumbs_down"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
