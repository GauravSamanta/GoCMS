package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gocms/connections"
	"gocms/models"
	"golang.org/x/crypto/bcrypt"
)

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

func ValidatePassword(userPassword string, bodyPassword string, c *gin.Context) {
	// take the password and hash it and compare it to the password in the database
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(bodyPassword))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}
}

func GetImage(c *gin.Context, id uint) models.Post {
	var image models.Post
	connections.DB.First(&image, id)
	return image
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

var LatestPost models.UserPostLink

func getLatestPostByTime(posts []models.UserPostLink) models.UserPostLink {
	for _, post := range posts {
		if post.Post.CreatedAt.After(LatestPost.Post.CreatedAt) {
			LatestPost = post
		}
	}

	return LatestPost
}

var SecondLatestPost models.UserPostLink

func getSecondLatestPostByTime(posts []models.UserPostLink, lastestPost models.UserPostLink) models.UserPostLink {
	for _, post := range posts {
		if post.Post.CreatedAt.After(SecondLatestPost.Post.CreatedAt) && post.Post.CreatedAt.Before(lastestPost.Post.CreatedAt) {
			SecondLatestPost = post
		}
	}
	return SecondLatestPost
}
