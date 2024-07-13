package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string 
	Email    string `json:"email" gorm:"unique,not null"` 
	Password string 
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
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`               
	Title         string         `json:"title" binding:"uppercase"`                        
	Description   string         `json:"description"`                                      
	Content       string         `json:"content"`                                          
	CreatedAt     time.Time      `json:"created_at" time_format:"2006-01-02"`              
	UpdatedAt     time.Time      `json:"updated_at" time_format:"2006-01-02"`              
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"` 
	FormattedDate string         `gorm:"-" json:"-"`
	//relationship with user
	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:UserID;" json:"-"`
}

type Image struct {
	UUID string `json:"uuid"`                               
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"` 
	Name string `json:"name"`                               
	Path string `json:"path"`                               
	Alt  string `json:"alt"`                                
	// PostID    uint           `json:"post_id"`                            
	// Post      Post           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time      `json:"created_at" time_format:"2006-01-02"`              
	UpdatedAt time.Time      `json:"updated_at" time_format:"2006-01-02"`              
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" time_format:"2006-01-02"` 
	//todo relationship with post
	// PostID uint `json:"post_id"`
	// Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:PostID;" json:"-"`
}
