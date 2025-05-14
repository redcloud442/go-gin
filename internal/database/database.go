package database

import (
	"fmt"
	"log"

	"github.com/inikhildubey/GoLang-Gin-boilerplate/config"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	cfg, err := config.LoadConfig() // Load database configs
	if err != nil {
		log.Fatal("❌ Failed to load configuration:", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")
	DB = database

	// Auto-migrate models
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Failed to auto-migrate database:", err)
	}
}
