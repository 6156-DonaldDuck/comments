package model

import "time"

type Comment struct {
	CommentId string `json:"comment_id"`
	AuthorId string `json:"author_id"`
	ArticleId string `json:"article_id"`
	Content string `json:"content"`
	CreateAt time.Time `json:"create_at"`
}