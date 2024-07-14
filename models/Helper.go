package models

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gorm.io/gorm"
)

func (p *Post) ThisHadTOBeAfterCreate(tx *gorm.DB) (err error) {
	// Assuming you have a method to get the current user and an image ID
	userID := GetCurrentUserID()

	link := UserPostLink{
		UserID: userID,
		PostID: p.ID,
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
	p.Content = string(MdToHTML([]byte(p.Content)))
	if len(p.Description) > 100 {
		p.Description = FirstN2(p.Description, 100, "....")
	}
}

func MdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
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
