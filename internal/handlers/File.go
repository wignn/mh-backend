package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/services"
	"github.com/wignn/mh-backend/pkg/utils"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "File is required")
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to open file")
			return
		}
		defer file.Close()

		filename, err := services.SaveImage(file, fileHeader.Filename)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to save image")
			return
		}

		utils.RespondJSON(c, http.StatusOK, gin.H{"url": filename}, "Image uploaded successfully")
	}
}

func GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Param("url")

		path, err := services.GetImage(url)
		if err != nil {
			utils.RespondJSON(c, http.StatusNotFound, nil, "File not found")
			return
		}

		c.File(path)
	}
}
