package controllers

import (
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"

	"github.com/gin-gonic/gin"
)

func AdminGetPosts(c *gin.Context) {
	var posts []models.UserPostImageLink

	result := connections.DB.Preload("User").Preload("Image").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for i := 0; i < len(posts); i++ {
		posts[i].Post.FormatAndTruncate()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts fetched successfully",
		"result":  posts,
	})

	// render.Render(c, http.StatusOK, postview.Posts(posts))
}

// create post
func AdminCreatePost(c *gin.Context) {
	var postbody struct {
		Title       string
		Description string
		Content     string
	}

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	post := models.Post{Title: postbody.Title, Description: postbody.Description, Content: postbody.Content, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result := connections.DB.Create(&post)

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
func AdminUpdatePost(c *gin.Context) {
	var postbody struct {
		ID          uint
		Title       string
		Description string
		Content     string
	}

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	result := connections.DB.Save(&models.Post{ID: postbody.ID, Title: postbody.Title, Description: postbody.Description, Content: postbody.Content, UpdatedAt: time.Now()})

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
func AdminDeletePost(c *gin.Context) {
	var postbody struct {
		ID uint
	}

	if c.Bind(&postbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
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
