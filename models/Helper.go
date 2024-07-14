package models

import (
	"strings"

	"gorm.io/gorm"
)

func (p *Post) ThisHadTOBeAfterCreate(tx *gorm.DB) (err error) {
	// Assuming you have a method to get the current user and an image ID
	userID := GetCurrentUserID()

	link := UserPostImageLink{
		UserID:  userID,
		PostID:  p.ID,
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
