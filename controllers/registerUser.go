package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"stock_market_simulation/m/constants"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
)

func GenerateLongUserID(id uint) string {
	return fmt.Sprintf("USR%06d", id) // Generating a simple example format
}

func ResgisterUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Username == "" || len(user.Username) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.UserNameValidationError,
		})
		return
	}
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.PasswordValidationError,
		})
		return
	}
	var existingUser models.Users
	result := initializers.DB.Where("username = ? OR user_id = ?", user.Username, user.UserID).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": constants.PasswordHashingError})
				return
			}
			user.Password = string(hashedPassword)
			initializers.DB.Create(&user)
			user.UserID = GenerateLongUserID(user.ID)
			initializers.DB.Save(&user)

			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.DatabaseError})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.UserAlreadyExists})
		return
	}

}
