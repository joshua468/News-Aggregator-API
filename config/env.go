package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
