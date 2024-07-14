package main

import (
	"fmt"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/Hrishikesh-Panigrahi/GoCMS/services"

	// "github.com/Hrishikesh-Panigrahi/GoCMS/services"

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
	r.GET("/register", controllers.RegisterPage)

	// authRoutes := r.Group("/auth/user")
	// authRoutes.POST("/login", controllers.Login)

	r.POST("/auth/user/login", controllers.Login)
	r.POST("/auth/user/register", controllers.Register)

	// for admin
	r.GET("/post", controllers.AdminGetPosts)
	r.POST("/post", controllers.AdminCreatePost)
	r.PUT("/post", controllers.AdminUpdatePost)
	r.DELETE("/post", controllers.AdminDeletePost)

	// for user
	r.GET("/user/post/", services.AuthMiddleware,controllers.GetPosts)
	r.GET("/user/post/:id", services.AuthMiddleware, controllers.GetPost)
	r.POST("/user/create/post", services.AuthMiddleware, controllers.CreatePost)
	r.PUT("/user/post/:id", services.AuthMiddleware, controllers.UpdatePost)
	r.DELETE("/user/post/:id", services.AuthMiddleware, controllers.DeletePost)

	r.GET("/contact", controllers.ContactUs)
	r.POST("/contactsucess", controllers.ContactMessage)

	r.Run()
}
