package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string         `json:"title" binding:"uppercase"`
	Description   string         `json:"description"`
	Content       string         `json:"content"`
	Category      string         `json:"category"`
	Tags          string         `json:"tags"`
	Path          string         `json:"path"`
	Alt           string         `json:"alt"`
	CreatedAt     time.Time      `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt     time.Time      `json:"updated_at" time_format:"2006-01-02"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"`
	FormattedDate string         `gorm:"-" json:"-"`
}

type UserPostLink struct {
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:UserID;" json:"-"`
	UserID uint `json:"user_id"`
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:PostID;" json:"-"`
	PostID uint `json:"post_id"`
}