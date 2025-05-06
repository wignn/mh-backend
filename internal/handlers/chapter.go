package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"github.com/wignn/mh-backend/pkg/utils" 
	"gorm.io/gorm"
)

func CreateChapter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var chapter model.Chapter

		if err := c.ShouldBindJSON(&chapter).Error; err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}

		res, err := services.CreateChapter(db, &chapter)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create chapter")
			return
		}

		utils.RespondJSON(c, http.StatusOK, res, "Chapter created successfully")
	}
}

func GetChapterById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Error getting ID")
			return
		}

		res, err := services.GetChapterById(db, id)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Internal server error")
			return
		}

		utils.RespondJSON(c, http.StatusOK, res, "Chapter retrieved successfully")
	}
}

func GetChapterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookIDStr := c.Param("bookID")
		bookID, err := strconv.Atoi(bookIDStr)
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, fmt.Sprintf("Invalid bookID '%s'", bookIDStr))
			return
		}

		chapter, err := services.GetChapterById(db, bookID)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to get chapter")
			return
		}

		utils.RespondJSON(c, http.StatusOK, chapter, "Chapter retrieved successfully")
	}
}

func UpdateChapter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Error getting ID")
			return
		}

		var chapter model.Chapter
		chapter.ID = uint(id)

		if err := c.ShouldBindJSON(&chapter).Error; err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}

		res, err := services.UpdateChapter(db, &chapter)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to update chapter")
			return
		}

		utils.RespondJSON(c, http.StatusOK, res, "Chapter updated successfully")
	}
}
