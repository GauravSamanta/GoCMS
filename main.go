package main

import (
	"fmt"

	"gocms/connections"
	"gocms/controllers"
	"gocms/routes"

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
	r.Static("/media", "./media")

	// backend routes
	router := r.Group("/")
	routes.AddRoutes(router)

	// sort these routes as well
	r.GET("/contact", controllers.ContactUs)
	r.POST("/contactsucess", controllers.ContactMessage)

	r.Run()
}
