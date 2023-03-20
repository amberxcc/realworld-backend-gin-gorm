package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	
	UserID      uint
	TagList     []Tag     `json:"tagList"`
	Comments    []Comment `json:"comments"`
}
