package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	view404 "github.com/Hrishikesh-Panigrahi/GoCMS/templates/404"
	postview "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Admin"

	"github.com/gin-gonic/gin"
)

func AdminGetPosts(c *gin.Context) {
	// get the current logged in user
	userID, _ := c.Get("userID")
	var user models.User
	connections.DB.First(&user, userID)

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

	latestPost := getLatestPostByTime(posts)
	secondLastestPost := getSecondLatestPostByTime(posts, latestPost)

	for i := 0; i < len(posts); i++ {
		for j := i + 1; j < len(posts); j++ {
			if posts[i].Post.CreatedAt.Unix() < posts[j].Post.CreatedAt.Unix() {
				posts[i], posts[j] = posts[j], posts[i]
			}
		}
	}

	var updatedPosts []models.UserPostLink
	for i := 0; i < len(posts); i++ {
		//remove the latest post and secondLastestPost from the list
		if posts[i].Post.ID != latestPost.Post.ID && posts[i].Post.ID != secondLastestPost.Post.ID {
			updatedPosts = append(updatedPosts, posts[i])
		}
	}

	render.Render(c, http.StatusOK, postview.AllPosts(updatedPosts, user))
}

func AdminGetPost(c *gin.Context) {
	if c.Request.Method == "GET" {
		var post models.Post

		id := c.Param("post_id")

		result := connections.DB.First(&post, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, view404.Page404("Post not found"))

		}
		content := string(MdToHTML([]byte(post.Content)))

		var LinkUserComments []models.LinkUserPostComment
		connections.DB.Preload("Comment").Preload("User").Where("post_id = ?", id).Find(&LinkUserComments)

		render.Render(c, http.StatusOK, postview.UserSinglepost(post, content, strconv.Itoa(len(LinkUserComments)), LinkUserComments))
	}

	if c.Request.Method == "POST" {
		CreateComment(c)
	}
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

}

// update post
func AdminUpdatePost(c *gin.Context) {
	ID := StringToUint(c.Param("id"))
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

	result := connections.DB.Save(&models.Post{ID: ID, Title: postbody.Title, Description: postbody.Description, Content: postbody.Content, UpdatedAt: time.Now()})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while updating the post",
		})
		return
	}

}

// delete post
func AdminDeletePost(c *gin.Context) {
	ID := StringToUint(c.Param("id"))

	result := connections.DB.Delete(&models.Post{}, ID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while deleting the post",
		})
		return
	}
}
