package models

import (
	"github.com/jinzhu/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Email string
}
