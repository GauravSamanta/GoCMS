package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `json:"email" gorm:"unique,not null"`
	UserName       string `json:"user_name" gorm:"unique,not null"`
	Password       string
	ProfileImgPath string   `json:"profile_img_path"`
	Alt            string   `json:"alt"`
	LocationID     uint     `gorm:"not null;DEFAULT:2" json:"location_id"`
	Location       Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:LocationID;" json:"location"`
	Bio            string   `json:"bio"`
	RoleID         uint     `gorm:"not null;DEFAULT:2" json:"role_id"`
	Role           Role     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:RoleID;" json:"-"`
}
