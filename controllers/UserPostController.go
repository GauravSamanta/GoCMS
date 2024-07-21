package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	view404 "github.com/Hrishikesh-Panigrahi/GoCMS/templates/404"
	postview "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Posts"
	processedviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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

		render.Render(c, http.StatusOK, postview.Singlepost(post.Title, post.Description, content, strconv.Itoa(len(LinkUserComments)), LinkUserComments, int(post.ID)))
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
		image_path := PostImageUpload(c)

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

}
