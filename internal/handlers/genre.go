package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/services"
	"gorm.io/gorm"
)

func CreateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre struct {
			Title string `json:"title"`
		}

		if err := c.ShouldBindJSON(&genre); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}
		createdGenre, err := services.CreateGenre(db, genre.Title)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create genre"})
			return
		}

		c.JSON(200, gin.H{"genre": createdGenre})
	}
}

func GetGenreById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		genre, err := services.GetGenreById(db, id)

		if err != nil {
			c.JSON(404, gin.H{"error": "Genre not found"})
			return
		}

		c.JSON(200, gin.H{"genre": genre})
	}
}
