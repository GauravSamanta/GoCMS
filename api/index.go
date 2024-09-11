package handler

import (
	"fmt"
	"net/http"

	"gocms/connections"
	"gocms/controllers"
	"gocms/routes"

	"github.com/gin-gonic/gin"
)

var print = fmt.Println
var (
	app *gin.Engine
)

func init() {

	connections.Connection()
	gin.SetMode(gin.ReleaseMode)
	connections.SyncDB()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := gin.Default()
	app.MaxMultipartMemory = 1
	print("Server is running on port 8080")
	app.Static("/static", "./static")
	app.Static("/media", "./media")

	// backend routes
	router := app.Group("/")
	routes.AddRoutes(router)

	// sort these routes as well
	app.GET("/contact", controllers.ContactUs)
	app.POST("/contactsucess", controllers.ContactMessage)

	app.ServeHTTP(w, r)
}
