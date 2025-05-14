package repositories

import (
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/database"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/models"
)

// Fetch users from DB (Mock Example)
func FetchUsers() ([]models.User, error) {
	users := []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	return users, nil
}

// GetUsers fetches all users from the database
// func GetUsers() ([]models.User, error) {
// 	var users []models.User
// 	result := database.DB.Find(&users) // Query the database

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return users, nil
// }

// GetUsers fetches all users if no ID is provided, otherwise fetches a single user by ID.
func GetUsers(id *uint) ([]models.User, error) {
	var users []models.User

	if id != nil { // If an ID is provided, fetch a single user
		var user models.User
		result := database.DB.First(&user, *id)

		if result.Error != nil {
			return nil, result.Error // Return error if user not found
		}
		return []models.User{user}, nil // Return a slice with a single user
	}

	// Fetch all users if no ID is provided
	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}
