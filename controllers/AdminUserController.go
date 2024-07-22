package controllers

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	view404 "github.com/Hrishikesh-Panigrahi/GoCMS/templates/404"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Admin"
	processedviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, views.CreateUser())
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
		var profileImagePath string

		if file != nil {
			profileImagePath = ProfileImageUpload(c)
		}

		var userbody struct {
			Name       string
			Email      string
			Password   string
			Role_ID    uint
			Username   string
			LocationID uint
		}
		if c.Bind(&userbody) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}

		user := models.User{
			Name:           userbody.Name,
			Email:          userbody.Email,
			Password:       userbody.Password,
			RoleID:         userbody.Role_ID,
			UserName:       userbody.Username,
			LocationID:     userbody.LocationID,
			ProfileImgPath: profileImagePath,
		}

		result := connections.DB.Create(&user)

		if result.Error != nil {
			render.Render(c, http.StatusInternalServerError, processedviews.Failure("Error while creating the user", result.Error.Error()))
		}
		render.Render(c, http.StatusInternalServerError, processedviews.Success("User Created Successfully", "", "", ""))
	}

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := connections.DB.First(&user, id)
	if result.Error != nil {
		render.Render(c, http.StatusNotFound, processedviews.Failure("Couldn't Find the user", result.Error.Error()))
		return
	}
	name := user.Name
	result = connections.DB.Unscoped().Delete(&user, id)

	if result.Error != nil {
		render.Render(c, http.StatusNotFound, processedviews.Failure("Couldn't delete the user", result.Error.Error()))
		return
	}

	render.Render(c, http.StatusOK, processedviews.Success("User Deleted Successfully", name, "", ""))
}

func UpdateUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Param("id")
		var user models.User
		result := connections.DB.First(&user, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, view404.Page404("User not found"))
			return
		}
		render.Render(c, http.StatusOK, views.EditUser(user))
		return
	}

	if c.Request.Method == "PUT" {
		id := c.Param("id")
		var user models.User
		result := connections.DB.First(&user, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, processedviews.Failure("Couldn't Find the user", result.Error.Error()))
			return
		}
		name := c.Request.FormValue("name")
		email := c.Request.FormValue("email")
		Role_ID := c.Request.FormValue("RoleID")

		fmt.Print(user.Email)

		fmt.Print(name, email, Role_ID)

		user.Name = name
		user.Email = email
		user.RoleID = StringToUint(Role_ID)
		connections.DB.Save(&user)

		render.Render(c, http.StatusOK, processedviews.Success("User Updated Successfully", "", "", ""))
		return
	}

}

func UpdatePassword(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, views.UpdatePassword())
		return
	}

	id := c.Param("id")
	var user models.User
	result := connections.DB.First(&user, id)

	if result.Error != nil {
		render.Render(c, http.StatusNotFound, processedviews.Failure("Couldn't Find the user", result.Error.Error()))
		return
	}

	password := c.Request.FormValue("password")

	user.Password = password
	result = connections.DB.Save(&user)
	if result.Error != nil {
		render.Render(c, http.StatusInternalServerError, processedviews.Failure("Couldn't Save the password", result.Error.Error()))
		return
	}
	render.Render(c, http.StatusOK, processedviews.Success("Password Updated Successfully", "", "", ""))
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := connections.DB.First(&user, id)
	if result.Error != nil {
		render.Render(c, http.StatusNotFound, view404.Page404("User not found"))
		return
	}

	var userPost models.UserPostLink

	connections.DB.Preload("Post").Where("user_id = ?", id).Find(&userPost)
	connections.DB.Preload("Location").Preload("Role").Find(&user)
	render.Render(c, http.StatusOK, views.GetUser(user, userPost))
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
	result = connections.DB.Preload("Role").Find(&users)
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
