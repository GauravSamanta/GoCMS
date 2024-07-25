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
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/User"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetProfile(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	connections.DB.First(&user, id)

	connections.DB.Model(&user).Preload("Role").Preload("Location").Find(&user)

	//get all the post from the user
	var userPostLink []models.UserPostLink

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPostLink)

	var userPostComment []models.LinkUserPostComment

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPostComment)

	render.Render(c, http.StatusOK, views.UserProfile(user, strconv.Itoa(len(userPostLink)), strconv.Itoa(len(userPostComment))))
}

func GetProfilePosts(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	connections.DB.First(&user, id)
	var userPostLink []models.UserPostLink

	for i := 0; i < len(userPostLink); i++ {
		userPostLink[i].Post.FormatAndTruncate()
	}

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPostLink)
	render.Render(c, http.StatusOK, views.UserProfilePosts(user, userPostLink))
}

// update post
func UpdatePost(c *gin.Context) {

	if c.Request.Method == "GET" {
		id := c.Param("id")
		var post models.Post

		result := connections.DB.First(&post, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, view404.Page404("Post not found"))
		}

		post.Content = string(MdToHTML([]byte(post.Content)))

		render.Render(c, http.StatusOK, views.UserEditPost(post))

	}
	if c.Request.Method == "POST" {

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
