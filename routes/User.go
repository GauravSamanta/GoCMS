package routes

import (
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/Hrishikesh-Panigrahi/GoCMS/services"
	"github.com/gin-gonic/gin"
)

func userRoutes(superRoute *gin.RouterGroup) {
	UserRoutes := superRoute.Group("api/v1/user")
	{
		userPostRoutes := UserRoutes.Group("/post")
		{
			userPostRoutes.POST("/create", services.AuthMiddleware, controllers.CreatePost)
			userPostRoutes.PUT("/:id", services.AuthMiddleware, controllers.UpdatePost)
			userPostRoutes.DELETE("/:id", services.AuthMiddleware, controllers.DeletePost)
		}
	}
	UserFormRoutes := superRoute.Group("/user")
	{
		userPostFormRoutes := UserFormRoutes.Group("/post")
		{
			userPostFormRoutes.GET("/create", services.AuthMiddleware, controllers.CreatePost)
			userPostFormRoutes.GET("/update/:id", services.AuthMiddleware, controllers.UpdatePost)
			userPostFormRoutes.GET("/delete/:id", services.AuthMiddleware, controllers.DeletePost)
		}
	}
}
