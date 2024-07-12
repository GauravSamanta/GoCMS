package services

import (
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/fossoreslp/go-uuid-v4"
	"github.com/gin-gonic/gin"
)

func AddImage(c *gin.Context, image_path string, uuid uuid.UUID) {
	//Todo get the user using the AwtToken? or maybe some other way

	var imagebody struct {
		Name string
		Alt  string
	}

	if c.Bind(&imagebody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	image := models.Image{UUID: uuid.String(), Name: imagebody.Name, Path: image_path, Alt: imagebody.Alt, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result := connections.DB.Create(&image)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating the image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "image created successfully",
		"result":  image,
	})
}

func DeleteImage(c *gin.Context, id string) {
	// var imagebody struct {
	// 	ID uint
	// }
	// if c.Bind(&imagebody) != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Invalid input",
	// 	})
	// 	return
	// }
	result := connections.DB.Delete(&models.Image{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while deleting the image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "image deleted successfully",
	})

}

func GetImage(c *gin.Context, id string) models.Image {
	var image models.Image
	connections.DB.First(&image, id)
	return image
}
