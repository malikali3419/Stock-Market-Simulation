package main

import (
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToTestDatabase()
}
func main() {
	initializers.DB.AutoMigrate(&models.StockData{})
}
