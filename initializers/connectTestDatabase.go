package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var TDB *gorm.DB

func ConnectToTestDatabase() {
	var err error
	dsn := os.Getenv("DB_TEST_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to Connect Database !")
	} else {
		log.Printf("Connection Success !")
	}
}
