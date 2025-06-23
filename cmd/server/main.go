package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/config"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/database"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/middleware"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/routes"
)

func main() {
	// Connect to Database
	database.InitDB()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	router := gin.Default()
	router.Use(middleware.RateLimiter())
	// Apply Middleware
	router.Use(middleware.Logger())

	// Register Routes
	routes.RegisterRoutes(router)

	// Start Server
	router.Run(":" + cfg.PORT)
}
