package config

import (
	"log"

	"github.com/joho/godotenv"
)

// trigger
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
