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
			users.POST("/", handlers.CreateUser(db))
			users.GET("/:id", handlers.GetUserByID(db))
			users.PUT("/:id", handlers.UpdateUser(db))
		}

		books := apiV1.Group("/books") 
		{
			books.POST("/", handlers.CreateBook(db))
			books.GET("/", handlers.GetBookByQuery(db))
			books.GET("/:id", handlers.GetBookByID(db))
			books.PUT("/:id", handlers.UpdateBook(db))
		}
		
	}
}
