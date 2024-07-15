package controllers

import (
	"net/http"
	"strconv"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(userPassword string, bodyPassword string, c *gin.Context) {
	// take the password and hash it and compare it to the password in the database
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(bodyPassword))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}
}

func GetImage(c *gin.Context, id uint) models.Post {
	var image models.Post
	connections.DB.First(&image, id)
	return image
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
