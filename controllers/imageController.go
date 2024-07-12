package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	"github.com/Hrishikesh-Panigrahi/GoCMS/services"

	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/index"
	"github.com/fossoreslp/go-uuid-v4"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ImageController(c *gin.Context) {
	render.Render(c, http.StatusOK, views.Index())
}

func ImageUpload(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		log.Error().Msgf("could not create multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file_array := form.File["file"]

	if len(file_array) == 0 || file_array[0] == nil {
		log.Error().Msgf("could not get the file array: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no file provided for image upload",
		})
		return
	}

	file := file_array[0]

	uuid, err := uuid.New()
	if err != nil {
		log.Error().Msgf("could not create the UUID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)

	filename := fmt.Sprintf("%s%s", uuid.String(), ext)
	image_path := filepath.Join("./media", filename)

	err = c.SaveUploadedFile(file, image_path)

	if err != nil {
		log.Error().Msgf("could not save file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// End saving to filesystem
	c.JSON(http.StatusOK, gin.H{
		"Id":   uuid.String(),
		"path": image_path,
	})

	// Adding Image to the DB
	services.AddImage(c, image_path, uuid)
}

func DeleteImage(c *gin.Context) {
	id := c.Param("id")

	//find the image from the DB
	image := services.GetImage(c, id)

	//delete the file from the filesystem
	err := os.Remove(image.Path)
	if err != nil {
		log.Error().Msgf("could not delete the file: %v", err)
		// No return because we have to remove the database entry nonetheless.
	}

	services.DeleteImage(c, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted successfully",
	})
}
