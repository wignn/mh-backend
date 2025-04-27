package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest model.User

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		user, err := services.RegisterUser(db, &userRequest)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(200, user)
	}
}

func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		user, err := services.GetUserByID(db, &id)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, user)
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		var user *model.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		user, err = services.UpdateUser(db, &id, user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(200, user)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		err = services.DeleteUser(db, &id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	}
}
