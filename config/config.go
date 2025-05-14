package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds database configurations
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// LoadConfig reads configuration values from .env file
func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Println("⚠️ Warning: Could not load .env file, using system environment variables.")
	}

	return &Config{
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
	}, nil
}
