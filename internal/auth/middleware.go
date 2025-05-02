package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/pkg/utils"
)

func AuthMIddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}
		claims, err := utils.ValidationToken(&token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.ID)
		c.Set("isAdmin", claims.IsAdmin)
		c.Next()

	}

}