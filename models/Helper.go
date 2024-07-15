package models

import (
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error().Msgf("Error hashing password for user %v: %v", u.Email, err)
		return
	}

	u.Password = string(hash)

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
