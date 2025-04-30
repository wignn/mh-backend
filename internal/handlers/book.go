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
		
		var book []*model.Book
		if err != nil {
			c.JSON(400, gin.H{"error": "query"})
			return
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(400, gin.H{"error": "page"})
			return
		}


		books, err := services.GetBookByQuery(db, book, &query, &page, &limit)
		if err != nil {
			log.Println("Error fetching book:", err)
			c.JSON(500, gin.H{"message": "not found"})
			return
		}

		c.JSON(200, books)
	}
}
