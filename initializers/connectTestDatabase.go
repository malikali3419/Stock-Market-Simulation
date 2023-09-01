package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"stock_market_simulation/m/constants"
)

var TDB *gorm.DB

func ConnectToTestDatabase() {
	var err error
	dsn := os.Getenv("DB_TEST_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(constants.FailedToConnectDatabase)
	} else {
		log.Printf(constants.ConnectionSuccess)
	}
}
