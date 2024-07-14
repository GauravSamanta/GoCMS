package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/services"

	// "github.com/Hrishikesh-Panigrahi/GoCMS/render"
	// views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Contact"

	// views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/user"
	// "fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}

	//set cookies and jwt token
	services.JwtToken(c, user)
	c.Redirect(http.StatusSeeOther, "http://localhost:8080/user/post")

}

func Register(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Role_ID  uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while hashing the password",
		})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		RoleID:   body.Role_ID,
	}

	result := connections.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating the user",
		})
		return
	}
	// http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
	// c.Redirect(http.StatusSeeOther, "http://localhost:8080/")
}
