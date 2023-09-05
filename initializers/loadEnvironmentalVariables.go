package initializers

import "github.com/Valgard/godotenv"

func LoadEnviromentalVariables() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
	if err := dotenv.Load(".env", ".env"); err != nil {
		panic(err)
	}
}
