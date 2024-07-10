package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//login
	render.Render(c, http.StatusOK, views.Login())
}
