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
	
	r.GET("/", controllers.Home)

	// authRoutes := r.Group("/auth/user")
	// authRoutes.POST("/login", controllers.Login)

	r.GET("/auth/user/login", controllers.Login)

	r.GET("/post", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)

	r.GET("/contact", controllers.ContactUs)
	r.POST("/contactsucess", controllers.ContactMessage)

	r.Run()
}
