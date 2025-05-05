package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/auth"
	"github.com/wignn/mh-backend/internal/handlers"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/health", handlers.HealthCheck())
	/**
	this is route for api v1
	*/
	apiV1 := r.Group("/api/v1")
	{
		/**
		this is route without authentication middleware
		*/

		//file routes
		apiV1.POST("/file", handlers.UploadImage())
		apiV1.GET("/file/:url", handlers.GetImage())
		
		//user routes
		apiV1.POST("/login", handlers.LoginUser(db))
		apiV1.POST("/register", handlers.CreateUser(db))

		//genre routes
		apiV1.GET("/book", handlers.GetBookByQuery(db))
		apiV1.GET("/book/:id", handlers.GetBookByID(db))

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
			aunthenticated.PUT("/book/:id", handlers.UpdateBook(db))
			aunthenticated.DELETE("/book/:id", handlers.DeleteBook(db))

			//genre routes
			aunthenticated.POST("/genre", handlers.CreateGenre(db))
			aunthenticated.GET("/genre/:id", handlers.GetGenreById(db))
			aunthenticated.GET("/genre", handlers.GetGenres(db))

			//bookmark routes
			aunthenticated.POST("/bookmark", handlers.CreateBookmark(db))
			aunthenticated.GET("/bookmark/:userId", handlers.GetBookByUser(db))
			aunthenticated.GET("/bookmark/:id", handlers.GetBookmarkById(db))
			aunthenticated.GET("/bookmark/:userId/:bookId", handlers.IsBookmark(db))
			aunthenticated.DELETE("/bookmark/:id", handlers.DeleteBookmark(db))
		}
	}
}
