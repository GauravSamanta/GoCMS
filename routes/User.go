package routes

import (
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/Hrishikesh-Panigrahi/GoCMS/middleware"
	"github.com/gin-gonic/gin"
)

func userRoutes(superRoute *gin.RouterGroup) {
	UserRoutes := superRoute.Group("/api/v1/user")
	{
		userPostRoutes := UserRoutes.Group("/post", middleware.AuthMiddleware)
		{
			userPostRoutes.POST("/create", controllers.CreatePost)
			userPostRoutes.PUT("/:id", controllers.UpdatePost)
			userPostRoutes.DELETE("/:id", controllers.DeletePost)
			userPostRoutes.POST("/comment/:post_id", controllers.GetPost)
		}
	}
	UserFormRoutes := superRoute.Group("/user")
	{
		userPostFormRoutes := UserFormRoutes.Group("/post", middleware.AuthMiddleware)
		{
			userPostFormRoutes.GET("/", controllers.GetPosts)
			userPostFormRoutes.GET("/:id", controllers.GetPost)
			userPostFormRoutes.GET("/create", controllers.CreatePost)
			userPostFormRoutes.GET("/update/:id", controllers.UpdatePost)
			userPostFormRoutes.GET("/delete/:id", controllers.DeletePost)
		}

		userProfleRoutes := UserFormRoutes.Group("/profile", middleware.AuthMiddleware)
		{
			userProfleRoutes.GET("/:id", controllers.GetProfile)
			userProfleRoutes.GET("/:id/posts", controllers.GetProfilePosts)
		}
	}
}
