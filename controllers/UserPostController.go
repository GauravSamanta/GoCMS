package controllers

import (
	// "errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	postview "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Posts"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.UserPostLink

	postresult := connections.DB.Find(&posts)
	if postresult.Error != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "404 Posts not found",
			"error": "Posts not found",
		})
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
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "404 Post not found",
			"error": "Post not found",
		})
	}

	// post.FormattedDate = post.CreatedAt.Format("02 January 2006")
	// post.Title = strings.ToUpper(post.Title)

	// post.Content = string(models.MdToHTML([]byte(post.Content)))

	render.Render(c, http.StatusOK, postview.Singlepost(post.Title, post.Description, post.Content))
}

// create post
func CreatePost(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"result":  post,
	})

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
	var postbody struct {
		ID uint
	}

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	image := GetImage(c, postbody.ID)

	err := os.Remove(image.Path)
	if err != nil {
		log.Error().Msgf("could not delete the file: %v", err)
		// No return because we have to remove the database entry nonetheless.
	}

	result := connections.DB.Delete(&models.Post{}, postbody.ID)

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
