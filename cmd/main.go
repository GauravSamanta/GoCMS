package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var print = fmt.Println

func main() {
	r := gin.Default()
	print("Server is running on port 8080")
	r.Static("/static", "../ui/static")

	r.LoadHTMLFiles("../ui/index.html")


	// r.LoadHTMLGlob("../ui/*")
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.Run()
}
