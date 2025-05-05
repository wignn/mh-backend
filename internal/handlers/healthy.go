package handlers

import "github.com/gin-gonic/gin"

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"message": "The server is running smoothly.",
			"version": "1.0.0",
			"timestamp": c.Request.Header.Get("X-Timestamp"),
		})

	}
}