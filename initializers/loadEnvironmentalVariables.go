package initializers

import "github.com/Valgard/godotenv"

func LoadEnviromentalVariables() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
	// You can also load several files
	if err := dotenv.Load(".env", ".env"); err != nil {
		panic(err)
	}
}
