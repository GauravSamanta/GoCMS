package routes

import (
	"github.com/gin-gonic/gin"
	"gocms/controllers"
)

func authRoutes(superRoute *gin.RouterGroup) {
	authRoutes := superRoute.Group("api/v1/auth/user")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/register", controllers.Register)
	}
	authGETRoutes := superRoute.Group("")
	{
		authGETRoutes.GET("/login", controllers.Login)
		authGETRoutes.GET("/register", controllers.Register)
	}

}
