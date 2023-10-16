package dto

import "time"

type QuestionHighlight struct {
	ID          int       `json:"id"`
	Author      string    `json:"author"`
	AuthorID    int       `json:"author_id"`
	Title       string    `json:"title"`
	Question    string    `json:"question"`
	ThumbsRate  int       `json:"thumbs_rate"`
	ThumbsUp    int       `json:"thumbs_up"`
	ThumbsDown  int       `json:"thumbs_down"`
	AnswerCount int       `json:"answer_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
