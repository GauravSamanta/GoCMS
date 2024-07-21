package models

import (
	"time"

	"gorm.io/gorm"
)

type Comments struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Comment       string         `json:"comment"`
	CreatedAt     time.Time      `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt     time.Time      `json:"updated_at" time_format:"2006-01-02"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"`
	FormattedDate string         `gorm:"-" json:"-"`
}

type LinkUserPostComment struct {
	UserID    uint     `json:"user_id"`
	PostID    uint     `json:"post_id"`
	CommentID uint     `json:"comment_id"`
	Post      Post     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:PostID;" json:"-"`
	User      User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:UserID;" json:"-"`
	Comment   Comments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:CommentID;" json:"-"`
}
