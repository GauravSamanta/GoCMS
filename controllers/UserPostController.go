package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	view404 "github.com/Hrishikesh-Panigrahi/GoCMS/templates/404"
	postview "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Posts"
	processedviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

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
	var posts []models.UserPostLink

	postresult := connections.DB.Find(&posts)
	if postresult.Error != nil {
		render.Render(c, http.StatusInternalServerError, view404.Page404("Posts not found"))
	}

	result := connections.DB.Preload("User").Preload("Post").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for i := 0; i < len(posts); i++ {
		posts[i].Post.FormatAndTruncate()
	}
	render.Render(c, http.StatusOK, postview.Posts(posts))
}

func GetPost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	result := connections.DB.First(&post, id)
	if result.Error != nil {
		render.Render(c, http.StatusNotFound, view404.Page404("Post not found"))

	}

	content := string(mdToHTML([]byte(post.Content)))
	render.Render(c, http.StatusOK, postview.Singlepost(post.Title, post.Description, content))
}

// create post
func CreatePost(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, postview.CreatePost())
		return
	}

	image_path := ImageUpload(c)

	var postbody struct {
		Title       string
		Description string
		Content     string
		Category    string
		Tags        string
		Alt         string
	}

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	//get the user id from the token
	user, _ := c.Get("user")
	userobj := user.(models.User)

	fmt.Println(userobj.ID)

	post := models.Post{Title: postbody.Title, Description: postbody.Description, Content: postbody.Content,
		Category: postbody.Category, Tags: postbody.Tags, Path: image_path, Alt: postbody.Alt, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result := connections.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating the post",
		})
		return
	}

	result = connections.DB.Create(&models.UserPostLink{UserID: userobj.ID, PostID: post.ID})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating the post",
		})
		return
	}

	render.Render(c, http.StatusOK, processedviews.Success("Post Created Successfully", "", "", ""))

}

// update post
func UpdatePost(c *gin.Context) {
	var postbody struct {
		Title       string
		Description string
		Content     string
		Category    string
		Tags        string
	}

	ID := StringToUint(c.Param("id"))

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	result := connections.DB.Save(&models.Post{ID: ID, Title: postbody.Title, Description: postbody.Description,
		Content: postbody.Content, Category: postbody.Category, Tags: postbody.Tags,
		UpdatedAt: time.Now()})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while updating the post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
}

// delete post
func DeletePost(c *gin.Context) {
	ID := StringToUint(c.Param("id"))

	image := GetImage(c, ID)

	err := os.Remove(image.Path)
	if err != nil {
		log.Error().Msgf("could not delete the file: %v", err)
		// No return because we have to remove the database entry nonetheless.
	}

	result := connections.DB.Delete(&models.Post{}, ID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while deleting the post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post Deleted successfully",
	})
}
