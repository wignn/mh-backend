package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/services"
	"github.com/wignn/mh-backend/pkg/utils"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest model.RegisterRequest

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}

		user := model.User{
			Username: userRequest.Username,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		userPtr, err := services.RegisterUser(db, &user)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}

		userResponse := struct {
			Name      string `json:"name"`
			Email     string `json:"email"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		}{
			Name:      userPtr.Username,
			Email:     userPtr.Email,
			CreatedAt: userPtr.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: userPtr.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		c.JSON(200, userResponse)
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

		userResponse := struct {
			Username  string `json:"Username"`
			Email     string `json:"email"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		}{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		c.JSON(200, userResponse)
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		var user *model.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Validation error"})
			return
		}
		user.ID = uint(id)

		user, err = services.UpdateUser(db, user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user"})
			return
		}

		userResponse := struct {
			Username  string `json:"name"`
			Email     string `json:"email"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		}{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		c.JSON(200, userResponse)
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

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": " Validation error"})
			return
		}

		existingUser, err := services.LoginUser(db, &user)
		
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(user.Username, int(user.ID), user.IsAdmin)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(200, gin.H{
			"user": gin.H{
				"id":       existingUser.ID,
				"username": existingUser.Username,
				"email":    existingUser.Email,
				"is_admin": existingUser.IsAdmin,
			},
			"token": token,
		})
	}
}
