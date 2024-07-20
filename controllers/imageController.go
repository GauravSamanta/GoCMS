package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func PostImageUpload(c *gin.Context) string {

	form, err := c.MultipartForm()
	if err != nil {
		log.Error().Msgf("could not create multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	file_array := form.File["file"]

	if len(file_array) == 0 || file_array[0] == nil {
		log.Error().Msgf("could not get the file array: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no file provided for image upload",
		})
		return ""
	}

	file := file_array[0]

	filename := file.Filename
	image_path := filepath.Join("./media/post", filename)

	err = c.SaveUploadedFile(file, image_path)

	if err != nil {
		log.Error().Msgf("could not save file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}

	return filepath.Join("../../", image_path)
}

func ProfileImageUpload(c *gin.Context) string {
	form, err := c.MultipartForm()
	if err != nil {
		log.Error().Msgf("could not create multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	file_array := form.File["file"]

	if len(file_array) == 0 || file_array[0] == nil {
		log.Error().Msgf("could not get the file array: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no file provided for image upload",
		})
		return ""
	}

	file := file_array[0]

	filename := file.Filename
	image_path := filepath.Join("./media/userProfile/", filename)

	err = c.SaveUploadedFile(file, image_path)

	if err != nil {
		log.Error().Msgf("could not save file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}

	return filepath.Join("../../", image_path)
}
