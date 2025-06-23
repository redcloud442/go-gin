package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds database configurations
type Config struct {
	DBConnectionString string
	PORT               string
}

// LoadConfig reads configuration values from .env file
func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Println("⚠️ Warning: Could not load .env file, using system environment variables.")
	}

	return &Config{
		DBConnectionString: os.Getenv("NEON_CONNECTION_STRING"),
		PORT:               os.Getenv("PORT"),
	}, nil
}
