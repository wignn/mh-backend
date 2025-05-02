package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/services"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "File is required"})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open file"})
			return
		}
		defer file.Close()

		filename, err := services.SaveImage(file, fileHeader.Filename)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save image"})
			return
		}

		c.JSON(200, gin.H{
			"url": filename,
		})
	}
}


func GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Param("url")

		path, err := services.GetImage(url)
		if err != nil {
			c.JSON(404, gin.H{"error": "File not found"})
			return
		}

		c.File(path)
	}
}
