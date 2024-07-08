package main

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"

	"github.com/gin-gonic/gin"
)

var print = fmt.Println

func init() {
	connections.LoadEnvVariables()
	connections.Connection()
	connections.SyncDB()
}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 1
	print("Server is running on port 8080")
	r.Static("/static", "./static")
	// r.LoadHTMLFiles("./templates/index.html","./templates/Contact/failure.html")
	// r.LoadHTMLFiles("./templates/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)

	r.POST("/contactus", func(c *gin.Context) {
		c.Request.ParseForm()

		email := c.Request.FormValue("email")
		name := c.Request.FormValue("name")
		message := c.Request.FormValue("message")

		_, err := mail.ParseAddress(email)
		if err != nil {
			c.HTML(http.StatusOK, "failure.html", gin.H{
				"title": "failure",
				"error": err,
			})
			return
		}

		print("Name: ", name)
		print("Email: ", email)
		print("Message: ", message)
		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":   "success",
			"message": message,
			"email":   email,
			"name":    name,
		})
	})

	r.Run()
}
