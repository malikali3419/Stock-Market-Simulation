package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
}

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
