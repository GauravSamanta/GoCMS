package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	postviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Posts"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/index"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"strings"

	"github.com/gin-gonic/gin"
)

func firstN2(s string, n int, suffix string) string {
	r := []rune(s)
	if len(r) > n {
		return string(r[:n]) + suffix
	}
	return s
}

func mdToHTML(md []byte) []byte {
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

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result := connections.DB.Find(&posts)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "404 Posts not found",
			"error": "Posts not found",
		})
	}

	for i := 0; i < len(posts); i++ {
		posts[i].FormattedDate = posts[i].CreatedAt.Format("02 January 2006")
		posts[i].Title = strings.ToUpper(posts[i].Title)
		if len(posts[i].Description) > 100 {
			posts[i].Description = firstN2(posts[i].Description, 100, "....")

		}
	}
	render.Render(c, http.StatusOK, views.Index(posts))
}

func GetPost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	result := connections.DB.First(&post, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "404 Post not found",
			"error": "Post not found",
		})
	}

	post.FormattedDate = post.CreatedAt.Format("02 January 2006")
	post.Title = strings.ToUpper(post.Title)

	post.Content = string(mdToHTML([]byte(post.Content)))

	render.Render(c, http.StatusOK, postviews.Singlepost(post.Title, post.Content))
}
