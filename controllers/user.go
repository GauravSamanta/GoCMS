package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	// "github.com/Hrishikesh-Panigrahi/GoCMS/render"
	// views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Contact"

	// views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/user"
	// "fmt"

	"github.com/Hrishikesh-Panigrahi/GoCMS/models"

	// "golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//login
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	// find the user
	var user models.User
	connections.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	// take the password and hash it and compare it to the password in the database

	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Invalid password",
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})

	// render.Render(c, 200, views.Contact())
}
