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

func GetUser(c *gin.Context) {
	var username = c.Param("username")
	var user models.Users
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
