package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    uint
	User      User   `json:"author"`
	Body      string `json:"body"`
	ArticleId uint
	Article   Article
}
