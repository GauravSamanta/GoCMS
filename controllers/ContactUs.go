package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContactUs(c *gin.Context) {
	message := c.PostForm("message")
	email := c.PostForm("email")
	name := c.PostForm("name")

	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Message: ", message)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
