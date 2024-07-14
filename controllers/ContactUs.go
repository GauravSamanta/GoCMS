package controllers

import (
	"net/http"
	"net/mail"

	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Contact"
	"github.com/gin-gonic/gin"
)

func ContactUs(c *gin.Context) {
	render.Render(c, http.StatusOK, views.Contact())
}

func ContactMessage(c *gin.Context) {
	c.Request.ParseForm()

	email := c.Request.FormValue("email")
	name := c.Request.FormValue("name")
	message := c.Request.FormValue("message")

	_, err := mail.ParseAddress(email)
	if err != nil {
		render.Render(c, http.StatusBadRequest, views.Failure("Failure in Parsing Email Address", name, err.Error()))
		return
	}

	render.Render(c, http.StatusAccepted, views.Success("Sucess", name, email, message))
}
