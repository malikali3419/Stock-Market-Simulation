package controllers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"time"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
}

func GetAllUsers(c *gin.Context) {

	// Retrieve value from Redis
	var users []models.Users
	val, err := initializers.REDISCLIENT.Get(context.Background(), "users").Result()
	if err == redis.Nil {

		result := initializers.DB.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error,
			})
			return
		}
		usersJSON, errorsUser := json.Marshal(users)
		if errorsUser != nil {
			panic("Failed to serialize users to JSON: " + err.Error())
		}
		key := "users"
		value := usersJSON
		expiration := time.Hour

		err := initializers.REDISCLIENT.Set(context.Background(), key, value, expiration).Err()
		if err != nil {
			panic("Failed to set data in Redis: " + err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
		return

	} else if err != nil {
		c.String(500, "Error retrieving value from Redis")
		return
	} else {

		err = json.Unmarshal([]byte(val), &users)
		if err != nil {
			panic("Failed to decode JSON users array: " + err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}

}
