package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
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
	res := initializers.DB.Find(&allStocks)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stocks": allStocks,
	})

}

func GetOneStock(c *gin.Context) {
	var ticker = c.Param("ticker")
	var foundedStock models.StockData
	result := initializers.DB.Where("ticker = ?", ticker).First(&foundedStock)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stock": foundedStock,
	})

}
