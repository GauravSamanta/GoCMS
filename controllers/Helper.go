package controllers

import (
	"strconv"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context, id uint) models.Post {
	var image models.Post
	connections.DB.First(&image, id)
	return image
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
