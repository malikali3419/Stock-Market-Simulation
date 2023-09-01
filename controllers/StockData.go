package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"net/http"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"time"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
	initializers.ConnectToRedis()
}

func AddStocks(c *gin.Context) {
	var stocks models.StockData
	if err := c.ShouldBindJSON(&stocks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var existingStocks models.StockData
	result := initializers.DB.Where("ticker = ? ", stocks.Ticker).First(&existingStocks)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		add_stocks := initializers.DB.Create(&stocks)
		if add_stocks.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": add_stocks.Error.Error(),
			})

			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"stock": stocks,
			})
			return
		}
	} else {
		initializers.DB.Model(&existingStocks).Updates(models.StockData{Ticker: stocks.Ticker,
			OpenPrice:  stocks.OpenPrice,
			Low:        stocks.Low,
			High:       stocks.High,
			ClosePrice: stocks.ClosePrice,
			Volume:     stocks.Volume,
		})
		c.JSON(http.StatusOK, gin.H{
			"stock": existingStocks,
		})

	}

}

func GetAllStocks(c *gin.Context) {
	var allStocks []models.StockData

	val, err := initializers.REDISCLIENT.Get(context.Background(), "stocks").Result()
	if err == redis.Nil {

		res := initializers.DB.Find(&allStocks)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error,
			})
			return
		}
		stockJSON, errorsStock := json.Marshal(allStocks)
		if errorsStock != nil {
			panic("Failed to serialize users to JSON: " + err.Error())
		}
		key := "stocks"
		value := stockJSON
		expiration := time.Hour

		err := initializers.REDISCLIENT.Set(context.Background(), key, value, expiration).Err()
		if err != nil {
			panic("Failed to set data in Redis: " + err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"stocks": allStocks,
		})
		return

	} else if err != nil {
		c.String(500, "Error retrieving value from Redis")
		return
	} else {

		err = json.Unmarshal([]byte(val), &allStocks)
		if err != nil {
			panic("Failed to decode JSON users array: " + err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"stocks": allStocks,
		})
	}

}

func GetOneStock(c *gin.Context) {
	var ticker = c.Param("ticker")
	var foundedStock models.StockData
	key := "stocks_" + ticker

	val, err := initializers.REDISCLIENT.Get(context.Background(), key).Result()
	if err == redis.Nil {

		result := initializers.DB.Where("ticker = ?", ticker).First(&foundedStock)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error,
			})
			return
		}
		stockJSON, errorsstock := json.Marshal(foundedStock)
		if errorsstock != nil {
			panic("Failed to serialize users to JSON: " + err.Error())
		}
		key := "stocks_" + ticker
		value := stockJSON
		expiration := time.Hour

		err := initializers.REDISCLIENT.Set(context.Background(), key, value, expiration).Err()
		if err != nil {
			panic("Failed to set data in Redis: " + err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"stocks": foundedStock,
		})
		return

	} else if err != nil {
		c.String(500, "Error retrieving value from Redis")
		return
	} else {

		err = json.Unmarshal([]byte(val), &foundedStock)
		if err != nil {
			panic("Failed to decode JSON users array: " + err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"stocks": foundedStock,
		})
	}

}
