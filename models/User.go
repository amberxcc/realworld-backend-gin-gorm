package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Collections []Article `json:"collections" gorm:"many2many:user_favorites"`
	Articles    []Article `json:"articles"`
	Follower    []User    `json:"followers" gorm:"many2many:user_follower"`
}
