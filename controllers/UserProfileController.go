package controllers

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/User"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	UserName := c.Param("user_name")

	var user models.User
	connections.DB.Where("user_name = ?", UserName).First(&user)

	fmt.Println(user.Location.CityName)

	render.Render(c, http.StatusOK, views.UserProfile(user))
}
