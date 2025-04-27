package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wignn/mh-backend/internal/config"
	"github.com/wignn/mh-backend/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	
	config.ConnectDb()

	port := os.Getenv("PORT")

	r := gin.Default()
	routes.InitRoutes(r, config.DB)

	if err := r.Run(":" + port); err != nil {
		panic("Error starting server: " + err.Error())
	}

}
