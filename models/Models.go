package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string // A regular string field
	Email    string `json:"email" gorm:"unique,not null"` // A pointer to a string, allowing for null values
	Password string // A regular stringx`` field
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

type Post struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`               // Standard field for the primary key
	Title         string         `json:"title" binding:"uppercase"`                        // A regular string field
	Description   string         `json:"description"`                                      // A regular string field
	Content       string         `json:"content"`                                          // A regular string field
	CreatedAt     time.Time      `json:"created_at" time_format:"2006-01-02"`              // Standard field for the creation time
	UpdatedAt     time.Time      `json:"updated_at" time_format:"2006-01-02"`              // Standard field for the update time
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"` // Standard field for soft delete
	FormattedDate string         `gorm:"-" json:"-"`
}
