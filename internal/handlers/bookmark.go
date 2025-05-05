package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"gorm.io/gorm"
)

func CreateBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookmark model.Bookmark

		if err := c.ShouldBindJSON(&bookmark); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		newBookmark, err := services.CreateBookmark(db, &bookmark)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create bookmark"})
			return
		}

		c.JSON(200, gin.H{"bookmark": newBookmark})
	}
}

func GetBookByUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		bookmarks, err := services.GetBookmarksByUser(db, userId)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get bookmarks"})
			return
		}

		c.JSON(200, gin.H{"bookmarks": bookmarks})
	}
}

func GetBookmarkById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		bookmark, err := services.GetBookmarkById(db, id)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get bookmark"})
			return
		}

		c.JSON(200, gin.H{"bookmark": bookmark})
	}
}

func IsBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		bookId, err := strconv.Atoi(c.Param("bookId"))

		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid book ID"})
			return
		}
		data, err := services.IsBookmarked(db, userId, bookId)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to check bookmark"})
			return
		}

		c.JSON(200, gin.H{"isBookmarked": true,
			"bookmark": data})
	}
}

func DeleteBookmark(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid bookmark ID"})
			return
		}

		err = services.DeleteBookmark(db, id)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete bookmark"})
			return
		}

		c.JSON(200, gin.H{"message": "Bookmark deleted successfully"})
	}
}
