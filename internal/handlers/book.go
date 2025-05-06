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


func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid request body")
			return
		}

		createdBook, err := services.CreateBook(db, &book)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create book")
			return
		}

		utils.RespondJSON(c, http.StatusOK, createdBook, "Book created successfully")
	}
}

func GetBookByQuery(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("query", "")
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit <= 0 {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid limit value")
			return
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page <= 0 {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid page value")
			return
		}

		books, total, err := services.GetBookByQuery(db, query, page, limit)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to fetch books")
			return
		}

		c.JSON(http.StatusOK, gin.H{
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
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid book ID")
			return
		}

		book, err := services.GetBookByID(db, &id)
		if err != nil {
			utils.RespondJSON(c, http.StatusNotFound, nil, "Book not found")
			return
		}

		utils.RespondJSON(c, http.StatusOK, book, "Book retrieved successfully")
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid book ID")
			return
		}

		var book model.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid request body")
			return
		}

		book.ID = uint(id)

		updatedBook, err := services.UpdateBook(db, &book)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to update book")
			return
		}

		utils.RespondJSON(c, http.StatusOK, updatedBook, "Book updated successfully")
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid book ID")
			return
		}

		if err := services.DeleteBook(db, id); err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to delete book")
			return
		}

		utils.RespondJSON(c, http.StatusOK, nil, "Book deleted successfully")
	}
}

func CreateGenreBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookGenre model.BookGenre

		if err := c.ShouldBindJSON(&bookGenre); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid request body")
			return
		}

		created, err := services.CreateBookGenre(db, &bookGenre)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create book genre")
			return
		}

		utils.RespondJSON(c, http.StatusOK, created, "Book genre created successfully")
	}
}
