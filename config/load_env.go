package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}
}
