package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string // A regular string field
	Email    string `gorm:"unique, not null"` // A pointer to a string, allowing for null values
	Password string // A regular string field
	Role     string // A regular string field
}