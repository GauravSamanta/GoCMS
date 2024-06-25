package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string // A regular string field
	Email    string `gorm:"unique, not null"` // A pointer to a string, allowing for null values
	Password string // A regular string field
	Role     string // A regular string field
}

type Post struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey, autoIncrement"` // Standard field for the primary key
	Title     string // A regular string field
	Content   string // A regular string field
	CreatedAt time.Time
	Updated   int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
