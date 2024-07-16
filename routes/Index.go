package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	adminRoutes(superRoute)
	userRoutes(superRoute)
	authRoutes(superRoute)
}
