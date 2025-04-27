package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/handlers"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	
	apiV1 := r.Group("/api/v1") 
	{
		apiV1.POST("/users", handlers.CreateUser(db))
		apiV1.GET("/users/:id", handlers.GetUserByID(db))
		apiV1.PUT("/users/:id", handlers.UpdateUser(db))
	}
}
