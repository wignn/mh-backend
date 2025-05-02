package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/auth"
	"github.com/wignn/mh-backend/internal/handlers"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	
	/**
	this is route for api v1
	*/
	apiV1 := r.Group("/api/v1") 
	{
		/**
		this is route without authentication middleware
		*/
		apiV1.POST("/sigout", handlers.CreateUser(db))
		apiV1.POST("/file", handlers.UploadImage())
		apiV1.GET("/file/:url", handlers.GetImage())
		apiV1.POST("/login", handlers.LoginUser(db))

		/**
		this is route with authentication middleware
		*/
		aunthenticated := apiV1.Group("/")
		{	
			//user routes
			aunthenticated.Use(auth.AuthMIddleware())
			aunthenticated.GET("/user/:id", handlers.GetUserByID(db))
			aunthenticated.PUT("/user/:id", handlers.UpdateUser(db))
			aunthenticated.DELETE("/user/:id", handlers.DeleteUser(db))

			//book routes
			aunthenticated.POST("/book", handlers.CreateBook(db))
			aunthenticated.GET("/book", handlers.GetBookByQuery(db))
			aunthenticated.GET("/book/:id", handlers.GetBookByID(db))
			aunthenticated.PUT("/book/:id", handlers.UpdateBook(db))
			aunthenticated.DELETE("/book/:id", handlers.DeleteBook(db))
		
		
			//genre routes
			aunthenticated.POST("/genre", handlers.CreateGenre(db))
		}		
	}
}
