package handlers

import (
	"net/http"
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
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}

		user := model.User{
			Username: userRequest.Username,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		userPtr, err := services.RegisterUser(db, &user)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to create user")
			return
		}

		response := gin.H{
			"username":   userPtr.Username,
			"email":      userPtr.Email,
			"created_at": userPtr.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at": userPtr.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		utils.RespondJSON(c, http.StatusOK, response, "User created successfully")
	}
}

func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid user ID")
			return
		}

		user, err := services.GetUserByID(db, &id)
		if err != nil {
			utils.RespondJSON(c, http.StatusNotFound, nil, "User not found")
			return
		}

		response := gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at": user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		utils.RespondJSON(c, http.StatusOK, response, "User fetched successfully")
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid user ID")
			return
		}

		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}
		user.ID = uint(id)

		updatedUser, err := services.UpdateUser(db, &user)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to update user")
			return
		}

		response := gin.H{
			"username":   updatedUser.Username,
			"email":      updatedUser.Email,
			"created_at": updatedUser.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at": updatedUser.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		utils.RespondJSON(c, http.StatusOK, response, "User updated successfully")
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Invalid user ID")
			return
		}
		err = services.DeleteUser(db, &id)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to delete user")
			return
		}
		utils.RespondJSON(c, http.StatusOK, nil, "User deleted successfully")
	}
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq model.LoginRequest

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			utils.RespondJSON(c, http.StatusBadRequest, nil, "Validation error")
			return
		}

		user, err := services.LoginUser(db, &model.User{
			Username: loginReq.Username,
			Password: loginReq.Password,
		})
		if err != nil {
			utils.RespondJSON(c, http.StatusUnauthorized, nil, "Invalid credentials")
			return
		}

		token, err := utils.GenerateToken(user.Username, int(user.ID), user.IsAdmin)
		if err != nil {
			utils.RespondJSON(c, http.StatusInternalServerError, nil, "Failed to generate token")
			return
		}

		response := gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
				"is_admin": user.IsAdmin,
			},
			"token": token,
		}
		utils.RespondJSON(c, http.StatusOK, response, "Login successful")
	}
}
