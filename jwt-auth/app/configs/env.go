package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	return os.Getenv("SECRET")
}

func EnvMySqlURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loding the .env file")
	}

	return os.Getenv("MYSQLURI")
}
