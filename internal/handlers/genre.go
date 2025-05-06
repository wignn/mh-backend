package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/services"
	"github.com/wignn/mh-backend/pkg/utils"
	"gorm.io/gorm"
)

func CreateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title string `json:"title" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}

		createdGenre, err := services.CreateGenre(db, req.Title)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create genre")
			return
		}

		utils.RespondJSON(c, http.StatusOK, gin.H{"genre": createdGenre}, "Genre created successfully")
	}
}

func GetGenreById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid genre ID")
			return
		}

		genre, err := services.GetGenreById(db, id)
		if err != nil {
			utils.RespondJSON(c, http.StatusNotFound, nil, "Genre not found")
			return
		}

		utils.RespondJSON(c, http.StatusOK, gin.H{"genre": genre}, "Genre fetched successfully")
	}
}

func GetGenres(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, err := services.GetGenres(db)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to get genres")
			return
		}

		utils.RespondJSON(c, http.StatusOK, gin.H{"data": genres}, "Genres fetched successfully")
	}
}
