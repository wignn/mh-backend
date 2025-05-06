package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"message": "The server is running.",
			"version": "1.0.0",
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime": time.Since(time.Now()).String(),
		})
	}
}