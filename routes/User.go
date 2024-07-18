package routes

import (
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/Hrishikesh-Panigrahi/GoCMS/middleware"
	"github.com/gin-gonic/gin"
)

func userRoutes(superRoute *gin.RouterGroup) {
	UserRoutes := superRoute.Group("api/v1/user")
	{
		userPostRoutes := UserRoutes.Group("/post")
		{
			userPostRoutes.POST("/create", middleware.AuthMiddleware, controllers.CreatePost)
			userPostRoutes.PUT("/:id", middleware.AuthMiddleware, controllers.UpdatePost)
			userPostRoutes.DELETE("/:id", middleware.AuthMiddleware, controllers.DeletePost)
		}
	}
	UserFormRoutes := superRoute.Group("/user")
	{
		userPostFormRoutes := UserFormRoutes.Group("/post")
		{
			userPostFormRoutes.GET("/", middleware.AuthMiddleware, controllers.GetPosts)
			userPostFormRoutes.GET("/:id", middleware.AuthMiddleware, controllers.GetPost)
			userPostFormRoutes.GET("/create", middleware.AuthMiddleware, controllers.CreatePost)
			userPostFormRoutes.GET("/update/:id", middleware.AuthMiddleware, controllers.UpdatePost)
			userPostFormRoutes.GET("/delete/:id", middleware.AuthMiddleware, controllers.DeletePost)
		}
	}
}
