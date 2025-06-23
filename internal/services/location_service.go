package services

import (
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/models"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/repositories"
)

func GetAllRegions() ([]models.RegionTable, error) {
	return repositories.FetchRegions()
}

func GetSpecificRegion(regionID int) ([]models.RegionTable, error) {
	return repositories.FetchSpecificRegion(regionID)
}

func GetAllProvinces() ([]models.ProvinceTable, error) {
	return repositories.FetchProvinces()
}

func GetSpecificMunicipality(provinceID int) ([]models.MunicipalityTable, error) {
	return repositories.FetchSpecificMunicipality(provinceID)
}

func GetSpecificBarangay(municipalityID int) ([]models.BarangayTable, error) {
	return repositories.FetchSpecificBarangay(municipalityID)
}
