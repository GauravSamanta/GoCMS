package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Admin"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, views.CreateUser())
		return
	}

	var userbody struct {
		Name     string
		Email    string
		Password string
		Role_ID  uint
	}
	if c.Bind(&userbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	user := models.User{
		Email:    userbody.Email,
		Password: userbody.Password,
		RoleID:   userbody.Role_ID,
	}

	result := connections.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating the user",
		})
		return
	}

}

func DeleteUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {
}

func UpdatePassword(c *gin.Context) {
}

func GetUser(c *gin.Context) {
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := connections.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching the users",
		})
		return
	}
	result= connections.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while preloading the role",
		})
		return
	}
	
	render.Render(c, http.StatusOK, views.Users(users))
}

func GetUsersByRole(c *gin.Context) {
}
