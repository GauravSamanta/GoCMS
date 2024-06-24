package main

import (
	"fmt"
	"net/http"
	"net/mail"

	
	"github.com/gin-gonic/gin"
	"github.com/Hrishikesh-Panigrahi/GoCMS/Connections" // Add this line to import the package
	// "github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
)



var print = fmt.Println


func init() {
	connections.LoadEnvVariables()
	connections.Connections()
	
}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 1
	print("Server is running on port 8080")

	// r.LoadHTMLFiles("./templates/index.html", "./templates/Contact/success.html", "./templates/Contact/failure.html")
	// r.LoadHTMLFiles("./templates/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

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

	r.Static("/static", "./static")

	r.Run()
}
