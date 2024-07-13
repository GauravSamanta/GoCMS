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
	RoleID   uint   `gorm:"not null;DEFAULT:2" json:"role_id"`
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
	//relationship with user
	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:UserID;" json:"-"`
}

type Image struct {
	UUID string `json:"uuid"`                               // A regular string field
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Standard field for the primary key
	Name string `json:"name"`                               // A regular string field
	Path string `json:"path"`                               // A regular string field
	Alt  string `json:"alt"`                                // A regular string field
	// PostID    uint           `json:"post_id"`                            // A regular string field
	// Post      Post           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time      `json:"created_at" time_format:"2006-01-02"`              // Standard field for the creation time
	UpdatedAt time.Time      `json:"updated_at" time_format:"2006-01-02"`              // Standard field for the update time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"` // Standard field for soft delete
	//todo relationship with post
	// PostID uint `json:"post_id"`
	// Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:PostID;" json:"-"`
}
