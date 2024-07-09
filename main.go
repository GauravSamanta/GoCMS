package main

import (
	"fmt"

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

	r.GET("/contact", controllers.ContactUs)
	r.POST("/contactsucess", controllers.ContactMessage)

	r.Run()
}
