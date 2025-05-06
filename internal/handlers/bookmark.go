package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"github.com/wignn/mh-backend/pkg/utils"
	"gorm.io/gorm"
)

func CreateBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookmark model.Bookmark

		if err := c.ShouldBindJSON(&bookmark); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid request body")
			return
		}

		newBookmark, err := services.CreateBookmark(db, &bookmark)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create bookmark")
			return
		}

		utils.RespondJSON(c, http.StatusOK, newBookmark, "Bookmark created successfully")
	}
}

func GetBookByUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		bookmarks, err := services.GetBookmarksByUser(db, userId)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to get bookmarks")
			return
		}

		utils.RespondJSON(c, http.StatusOK, bookmarks, "Bookmarks retrieved successfully")
	}
}

func GetBookmarkById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		bookmark, err := services.GetBookmarkById(db, id)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to get bookmark")
			return
		}

		utils.RespondJSON(c, http.StatusOK, bookmark, "Bookmark retrieved successfully")
	}
}

func IsBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid user ID")
			return
		}

		bookId, err := strconv.Atoi(c.Param("bookId"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid book ID")
			return
		}

		data, err := services.IsBookmarked(db, userId, bookId)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to check bookmark")
			return
		}

		utils.RespondJSON(c, http.StatusOK, gin.H{"isBookmarked": true, "bookmark": data}, "Bookmark check completed")
	}
}

func DeleteBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid bookmark ID")
			return
		}

		err = services.DeleteBookmark(db, id)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to delete bookmark")
			return
		}

		utils.RespondJSON(c, http.StatusOK, nil, "Bookmark deleted successfully")
	}
}
