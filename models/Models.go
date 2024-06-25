package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string // A regular string field
	Email    string `json:"email" gorm:"unique,not null"` // A pointer to a string, allowing for null values
	Password string // A regular string field
	Role     string // A regular string field
}

type Post struct {
    ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"` // Standard field for the primary key
    Title     string         `json:"title" binding:"uppercase"`                              // A regular string field
    Content   string         `json:"content"`                            // A regular string field
    CreatedAt time.Time      `json:"created_at" time_format:"2006-01-02"` // Standard field for the creation time
    UpdatedAt time.Time      `json:"updated_at" time_format:"2006-01-02"` // Standard field for the update time
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"`            // Standard field for soft delete
	FormattedDate   string         `gorm:"-" json:"-"` 
}

