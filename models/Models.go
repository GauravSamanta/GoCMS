package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `json:"email" gorm:"unique,not null"`
	Password string
	RoleID   uint `gorm:"not null;DEFAULT:2" json:"role_id"`
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

type Post struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string         `json:"title" binding:"uppercase"`
	Description   string         `json:"description"`
	Content       string         `json:"content"`
	CreatedAt     time.Time      `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt     time.Time      `json:"updated_at" time_format:"2006-01-02"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"`
	FormattedDate string         `gorm:"-" json:"-"`
}

type Image struct {
	UUID      string         `json:"uuid"`
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Alt       string         `json:"alt"`
	CreatedAt time.Time      `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt time.Time      `json:"updated_at" time_format:"2006-01-02"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"`
}

type UserPostImageLink struct {
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:UserID;" json:"-"`
	UserID  uint  `json:"user_id"`
	Post    Post  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:PostID;" json:"-"`
	PostID  uint  `json:"post_id"`
	Image   Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:ImageID;" json:"-"`
	ImageID uint  `json:"image_id"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// Assuming you have a method to get the current user and an image ID
	userID := GetCurrentUserID()
	imageID := GetDefaultImageID()

	link := UserPostImageLink{
		UserID:  userID,
		PostID:  p.ID,
		ImageID: imageID,
	}

	if err := tx.Create(&link).Error; err != nil {
		return err
	}

	return nil
}

// function of the post class
func (p *Post) FormatAndTruncate() {
	p.FormattedDate = p.CreatedAt.Format("02 January 2006")
	p.Title = strings.ToUpper(p.Title)
	if len(p.Description) > 100 {
		p.Description = FirstN2(p.Description, 100, "....")
	}
}

func FirstN2(s string, n int, suffix string) string {
    if len(s) <= n {
        return s
    }
    return s[:n] + suffix
}

func GetCurrentUserID() uint {

	return 1
}

func GetDefaultImageID() uint {

	return 1
}
