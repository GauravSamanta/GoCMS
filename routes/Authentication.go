package routes

import (
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup) {
	authRoutes := superRoute.Group("/auth/user")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/register", controllers.Register)
	}

}
