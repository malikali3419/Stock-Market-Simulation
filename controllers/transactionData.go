package controllers

import (
	"errors"
	"fmt"
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

func GenerateTransaction(id uint) string {
	return fmt.Sprintf("TRANS%06d", id) // Generating a simple example format
}

func DoTransaction(c *gin.Context) {
	var transaction models.TransactionData
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var user models.Users
	err := initializers.DB.Where("id = ?", transaction.UserID).First(&user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	var stockData models.StockData
	result := initializers.DB.Where("ticker = ?", transaction.Ticker).First(&stockData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ticker not found!",
		})
		return
	}
	initializers.DB.Create(&transaction)
	if transaction.TransactionType == "buy" {
		transactionPrice := transaction.TransactionVolume * stockData.Low

		newTransaction := initializers.DB.Model(&transaction).Updates(models.TransactionData{
			TransactionPrice: transactionPrice,
			TransactionID:    GenerateTransaction(transaction.ID),
		})
		if int(user.Balance) < transactionPrice {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User has enough balance for this transaction !",
			})
			return
		}

		if newTransaction.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": newTransaction.Error.Error(),
			})
			return
		}
		updatedBalance := user.Balance - float32(transactionPrice)
		updatebalanceBuy := map[string]interface{}{
			"Balance": updatedBalance,
		}
		initializers.DB.Model(&user).Updates(updatebalanceBuy)
		c.JSON(http.StatusOK, gin.H{
			"transaction": transaction,
		})

	} else if transaction.TransactionType == "sell" {
		transactionPrice := transaction.TransactionVolume * stockData.High
		newTransaction := initializers.DB.Model(&transaction).Updates(models.TransactionData{
			TransactionPrice: transactionPrice,
			TransactionID:    GenerateTransaction(transaction.ID),
		})
		if newTransaction.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": newTransaction.Error.Error(),
			})
			return
		}
		updatedBalance := user.Balance + float32(transactionPrice)
		updatebalanceSell := map[string]interface{}{
			"Balance": updatedBalance,
		}
		initializers.DB.Model(&user).Updates(updatebalanceSell)
		c.JSON(http.StatusOK, gin.H{
			"transaction": transaction,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Transaction Type",
		})
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
			"error": "User not found",
		})
		return
	}
	var transactions []models.TransactionData
	initializers.DB.Where("user_id = ? AND created_at BETWEEN ? AND ?", user.ID, startTimeStamp, endTimeStamp).Preload("User").Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"transaction data": transactions,
	})
}
