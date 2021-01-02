package models

import "github.com/jinzhu/gorm"

// User is our app user model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
