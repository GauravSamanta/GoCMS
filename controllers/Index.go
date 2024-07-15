package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/index"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	render.Render(c, http.StatusOK, views.Index())
}

func RegisterPage(c *gin.Context) {
	render.Render(c, http.StatusOK, views.Registration())
}
