package models

import "github.com/jinzhu/gorm"

// Post struct
type Post struct {
	gorm.Model
	UserID uint
	Text   string `json:"text"`
}
