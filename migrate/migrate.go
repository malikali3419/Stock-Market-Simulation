package main

import (
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
}
func main() {
	initializers.DB.AutoMigrate(&models.TransactionData{})
}