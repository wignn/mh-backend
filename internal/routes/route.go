package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/handlers"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	
	apiV1 := r.Group("/api/v1") 
	{
		users := apiV1.Group("/users")
		{
			users.POST("/users", handlers.CreateUser(db))
			users.GET("/users/:id", handlers.GetUserByID(db))
			users.PUT("/users/:id", handlers.UpdateUser(db))
		}

		books := apiV1.Group("/books") 
		{
			books.POST("/books", handlers.CreateBook(db))
			books.GET("/books", handlers.GetBookByQuery(db))
		}
		
	}
}
