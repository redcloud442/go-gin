package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/database"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/middleware"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/routes"
)

func main() {
	// Connect to Database
	database.InitDB()

	router := gin.Default()

	// Apply Middleware
	router.Use(middleware.Logger())

	// Register Routes
	routes.RegisterRoutes(router)

	// Start Server
	router.Run(":8080")
}
