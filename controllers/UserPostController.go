package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gocms/connections"
	"gocms/models"
	"gocms/render"
	view404 "gocms/templates/404"
	postview "gocms/templates/Posts"
	processedviews "gocms/templates/Processed"
)

func GetPosts(c *gin.Context) {
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

	render.Render(c, http.StatusOK, postview.Posts(updatedPosts, user, latestPost, secondLastestPost))
}

func GetPost(c *gin.Context) {
	if c.Request.Method == "GET" {
		var post models.Post

		id := c.Param("id")

		result := connections.DB.First(&post, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, view404.Page404("Post not found"))

		}
		content := string(MdToHTML([]byte(post.Content)))

		var LinkUserComments []models.LinkUserPostComment
		connections.DB.Preload("Comment").Preload("User").Where("post_id = ?", id).Find(&LinkUserComments)

		render.Render(c, http.StatusOK, postview.Singlepost(post, content, strconv.Itoa(len(LinkUserComments)), LinkUserComments))
	}

	if c.Request.Method == "POST" {
		CreateComment(c)
	}

}

// create post
func CreatePost(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, postview.CreatePost())
		return
	}

	if c.Request.Method == "POST" {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}
		defer file.Close()
		var image_path string

		if file != nil {
			image_path = PostImageUpload(c)
		}

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

		userID, _ := c.Get("userID")
		var user models.User
		connections.DB.First(&user, userID)

		post := models.Post{Title: postbody.Title, Description: postbody.Description, Content: postbody.Content,
			Category: postbody.Category, Tags: postbody.Tags, Path: image_path, Alt: postbody.Alt, CreatedAt: time.Now(), UpdatedAt: time.Now()}

		result := connections.DB.Create(&post)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating the post",
			})
			return
		}

		result = connections.DB.Create(&models.UserPostLink{UserID: user.ID, PostID: post.ID})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating the post",
			})
			return
		}
		render.Render(c, http.StatusOK, processedviews.Success("Post Created Successfully", "", "", ""))
	}
}
