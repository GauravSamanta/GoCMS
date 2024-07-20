package controllers

import (
	"net/http"
	"strconv"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/User"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	connections.DB.First(&user, id)

	connections.DB.Model(&user).Preload("Role").Preload("Location").Find(&user)

	//get all the post from the user
	var userPostLink []models.UserPostLink

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPostLink)

	render.Render(c, http.StatusOK, views.UserProfile(user, strconv.Itoa(len(userPostLink))))
}

func GetProfilePosts(c *gin.Context){
	id := c.Param("id")

	var user models.User
	connections.DB.First(&user, id)
	var userPostLink []models.UserPostLink

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPostLink)
	render.Render(c, http.StatusOK, views.UserProfilePosts(user, userPostLink))
}