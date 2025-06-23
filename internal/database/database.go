package database

import (
	"fmt"
	"log"

	"github.com/inikhildubey/GoLang-Gin-boilerplate/config"

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

	database, err := gorm.Open(postgres.Open(cfg.DBConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")
	DB = database

	// // Auto-migrate models
	// err = database.AutoMigrate(&models.RegionTable{})
	// err = database.AutoMigrate(&models.ProvinceTable{})
	// err = database.AutoMigrate(&models.MunicipalityTable{})
	// err = database.AutoMigrate(&models.BarangayTable{})
	if err != nil {
		log.Fatal("❌ Failed to auto-migrate database:", err)
	}
}
