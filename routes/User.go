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

	UserGETRoutes := superRoute.Group("/:slug")
	{
		userPostGETRoutes := UserGETRoutes.Group("/post", middleware.AuthMiddleware)
		{
			userPostGETRoutes.GET("/", controllers.GetPosts)
			userPostGETRoutes.GET("/:id", controllers.GetPost)
			userPostGETRoutes.GET("/create", controllers.CreatePost)
			userPostGETRoutes.GET("/update/:id", controllers.UpdatePost)
			userPostGETRoutes.GET("/delete/:id", controllers.DeletePost)
		}

		userProfleRoutes := UserGETRoutes.Group("/profile", middleware.AuthMiddleware)
		{
			userProfleRoutes.GET("/:id", controllers.GetProfile)
			userProfleRoutes.GET("/:id/posts", controllers.GetProfilePosts)
		}
	}
}
