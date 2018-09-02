package domain

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	ArticleID int       `json:"articleId"`
	CreatedAt time.Time `json:"createdAt"`
}

type Comments []Comment
