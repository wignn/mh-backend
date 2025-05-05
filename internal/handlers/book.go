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

		c.JSON(200,
			gin.H{
				"data": createdBook,
			})
	}
}

func GetBookByQuery(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("query", "")

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit <= 0 {
			c.JSON(400, gin.H{"error": "invalid limit"})
			return
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page <= 0 {
			c.JSON(400, gin.H{"error": "invalid page"})
			return
		}

		books, total, err := services.GetBookByQuery(db, query, page, limit)
		if err != nil {
			log.Println("Error fetching book:", err)
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		c.JSON(200, gin.H{
			"data":  books,
			"page":  page,
			"limit": limit,
			"total": total,
		})
	}
}

func GetBookByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid book ID"})
			return
		}

		book, err := services.GetBookByID(db, &id)

		if err != nil {
			c.JSON(404, gin.H{"message": "not found"})
			return
		}

		c.JSON(200,
			gin.H{
				"data": book,
			})
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		var book model.Book

		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Book ID"})
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		book.ID = uint(id)

		updatedBook, err := services.UpdateBook(db, &book)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update book"})
			return
		}

		c.JSON(200, gin.H{
			"data": updatedBook,
		})
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid book ID"})
			return
		}

		if err := services.DeleteBook(db, id); err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete book"})
			return
		}

		c.JSON(200, gin.H{"message": "Book deleted successfully"})
	}
}

func CreateGenreBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var BookGenre model.BookGenre

		if err := c.ShouldBindJSON(&BookGenre); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		createdBookGenre, err := services.CreateBookGenre(db, &BookGenre)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create book genre"})
			return
		}

		c.JSON(200,
			gin.H{
				"data": createdBookGenre,
			})
	}
}
