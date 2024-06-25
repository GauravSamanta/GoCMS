package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"

	"github.com/gin-gonic/gin"
	"strings"
)

func firstN2(s string, n int, suffix string) string {
	r := []rune(s)
	if len(r) > n {
		return string(r[:n]) + suffix
	}
	return s
}

func GetPost(c *gin.Context) {
	var posts []models.Post

	result := connections.DB.Find(&posts)
	if result.Error != nil {
		return
	}

	for i := 0; i < len(posts); i++ {
		posts[i].FormattedDate = posts[i].CreatedAt.Format("02 January 2006")
		posts[i].Title = strings.ToUpper(posts[i].Title)
		if len(posts[i].Content) > 100 {
			posts[i].Content = firstN2(posts[i].Content, 100, "....")

		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
	})
}
