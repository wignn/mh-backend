package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var book model.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		createdBook, err := services.CreateBook(db, &book)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create book"})
			return
		}

		c.JSON(200, createdBook)
	}
}

func GetBookByQuery(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("query")
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(400, gin.H{"error": "query"})
			return
		}
		
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(400, gin.H{"error": "page"})
			return
		}

		log.Printf("Query: %s, Limit: %s, Page: %s\n", query, limit, page)
	}
}
