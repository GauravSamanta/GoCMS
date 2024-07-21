package models

import (
	"gorm.io/gorm"
)

type Location struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	CityName    string `gorm:"uniqueIndex:idx_first_second" json:"city_name"`
	StateName   string `gorm:"uniqueIndex:idx_first_second" json:"state_name"`
	CountryName string `gorm:"uniqueIndex:idx_first_second" json:"country_name"`
	ZipCode     string `gorm:"uniqueIndex:idx_first_second" json:"zip_code"`
}

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}


