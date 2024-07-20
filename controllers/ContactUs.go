package controllers

import (
	"net/http"
	"net/mail"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Contact"
	Processedviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gin-gonic/gin"
)

func ContactUs(c *gin.Context) {
	userID, _ := c.Get("userID")
	var user models.User
	connections.DB.Find(&user, userID)
	render.Render(c, http.StatusOK, views.Contact(user))
}

func ContactMessage(c *gin.Context) {
	c.Request.ParseForm()

	email := c.Request.FormValue("email")
	name := c.Request.FormValue("name")
	message := c.Request.FormValue("message")

	_, err := mail.ParseAddress(email)
	if err != nil {
		render.Render(c, http.StatusBadRequest, Processedviews.Failure("Failure in Parsing Email Address", err.Error()))
		return
	}

	render.Render(c, http.StatusAccepted, Processedviews.Success("Sucess", name, email, message))
}
