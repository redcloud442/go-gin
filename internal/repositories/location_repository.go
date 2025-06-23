package repositories

import (
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/database"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/models"
)

// Fetch users from DB (Mock Example)
func FetchRegions() ([]models.RegionTable, error) {
	var regions []models.RegionTable
	result := database.DB.Find(&regions)

	if result.Error != nil {
		return nil, result.Error
	}

	return regions, nil
}

func FetchProvinces() ([]models.ProvinceTable, error) {
	var provinces []models.ProvinceTable
	result := database.DB.Find(&provinces)

	if result.Error != nil {
		return nil, result.Error
	}

	return provinces, nil
}

func FetchSpecificRegion(regionID int) ([]models.RegionTable, error) {
	var region models.RegionTable
	result := database.DB.First(&region, regionID)

	if result.Error != nil {
		return nil, result.Error
	}

	return []models.RegionTable{region}, nil
}

func FetchSpecificMunicipality(provinceID int) ([]models.MunicipalityTable, error) {
	var municipalities []models.MunicipalityTable
	result := database.DB.Where("province_id = ?", provinceID).Find(&municipalities)

	if result.Error != nil {
		return nil, result.Error
	}

	return municipalities, nil
}

func FetchSpecificBarangay(municipalityID int) ([]models.BarangayTable, error) {
	var barangays []models.BarangayTable
	result := database.DB.Where("municipality_id = ?", municipalityID).Find(&barangays)

	if result.Error != nil {
		return nil, result.Error
	}

	return barangays, nil
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
