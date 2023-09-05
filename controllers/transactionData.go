package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stock_market_simulation/m/constants"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"time"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
	initializers.ConnectToRedis()
}

func GenerateTransaction(id uint) string {
	return fmt.Sprintf("TRANS%06d", id)
}

func DoTransaction(task TransactionTask) {

	c := task.Context
	transaction := task.Transaction

	var user models.Users
	err := initializers.DB.Where("id = ?", transaction.UserID).First(&user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.UserNotFound,
		})
		return
	}
	var stockData models.StockData
	result := initializers.DB.Where("ticker = ?", transaction.Ticker).First(&stockData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.TickerNotFound,
		})
		return
	}
	initializers.DB.Create(&transaction)
	if transaction.TransactionType == "buy" {
		transactionPrice := float32(transaction.TransactionVolume) * stockData.Low
		if stockData.Volume < transaction.TransactionVolume {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.NotEnoughStocks,
			})
			return
		}
		initializers.DB.Model(&transaction).Updates(models.TransactionData{
			TransactionPrice: transactionPrice,
			TransactionID:    GenerateTransaction(transaction.ID),
		})
		if user.Balance < transactionPrice {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constants.UserHasNotEnoughBalance,
			})
			return
		}
		updateVolume := stockData.Volume - transaction.TransactionVolume
		initializers.DB.Model(&stockData).Updates(models.StockData{
			Volume: updateVolume,
		})
		updatedBalance := user.Balance - float32(transactionPrice)
		updateBalanceBuy := map[string]interface{}{
			"Balance": updatedBalance,
		}
		initializers.DB.Model(&user).Updates(updateBalanceBuy)
	} else if transaction.TransactionType == "sell" {
		transactionPrice := float32(transaction.TransactionVolume) * stockData.High
		initializers.DB.Model(&transaction).Updates(models.TransactionData{
			TransactionPrice: transactionPrice,
			TransactionID:    GenerateTransaction(transaction.ID),
		})
		updateVolume := stockData.Volume + transaction.TransactionVolume
		initializers.DB.Model(&stockData).Updates(models.StockData{
			Volume: updateVolume,
		})
		updatedBalance := user.Balance + float32(transactionPrice)
		updateBalanceSell := map[string]interface{}{
			"Balance": updatedBalance,
		}
		initializers.DB.Model(&user).Updates(updateBalanceSell)
	}
}

func GetAllTransactionOFUser(c *gin.Context) {
	var userId = c.Param("user_id")
	var user models.Users
	err := initializers.DB.Where("id = ?", userId).First(&user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	var transactions []models.TransactionData
	initializers.DB.Where("user_id = ?", user.ID).Preload("User").Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"transaction data": transactions,
	})
}

func GetTransactionDataBetweenTime(c *gin.Context) {
	var userId = c.Param("user_id")
	var startTimeStamp = c.Param("start_time")
	var endTimeStamp = c.Param("end_time")
	var user models.Users
	err := initializers.DB.Where("id = ?", userId).First(&user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err2": constants.UserNotFound,
		})
		return
	}
	inputLayout := constants.DateInputLayout
	startDateFormated, err2 := time.Parse(inputLayout, startTimeStamp)
	if err2 != nil {
		fmt.Println(constants.ErrorInParsingDateToStr, err)
		return
	}
	endDateFormated, err3 := time.Parse(inputLayout, endTimeStamp)
	if err3 != nil {
		fmt.Println(constants.ErrorInParsingDateToStr, err)
		return
	}
	outputLayout := constants.DateOutPutLayout
	outputDateStrStartDate := startDateFormated.Format(outputLayout)
	outputDateStrEndDate := endDateFormated.Format(outputLayout)
	var transactions []models.TransactionData
	initializers.DB.Where("user_id = ? AND created_at BETWEEN ? AND ?", user.ID, outputDateStrStartDate, outputDateStrEndDate).Preload("User").Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"transaction data": transactions,
	})
}
