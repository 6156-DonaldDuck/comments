package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	AuthorId uint `json:"author_id"`
	ArticleId uint `json:"article_id"`
	Content string `json:"content"`
}