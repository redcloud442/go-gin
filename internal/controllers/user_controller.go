package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/models"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/repositories"
)

// Get all users
func GetUsers(c *gin.Context) {
	users, err := repositories.FetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func FetchUsers(c *gin.Context) {
	idParam := c.Query("id") // Get ID from query parameters
	var id *uint

	if idParam != "" { // Convert ID to uint if provided
		parsedID, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		uid := uint(parsedID)
		id = &uid
	}

	users, err := repositories.GetUsers(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUserController handles the user creation request
func CreateUserController(c *gin.Context) {
	var user models.User

	// Bind JSON request to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert user into DB
	if err := repositories.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
