package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/services"
)

// Get all users
func GetAllRegions(c *gin.Context) {
	regions, err := services.GetAllRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)
}

func GetAllProvinces(c *gin.Context) {
	provinces, err := services.GetAllProvinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, provinces)
}

func GetSpecificRegion(c *gin.Context) {
	regionID, err := strconv.Atoi(c.Param("region_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
		return
	}
	region, err := services.GetSpecificRegion(regionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

func GetSpecificProvince(c *gin.Context) {
	provinceID, err := strconv.Atoi(c.Param("province_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID"})
		return
	}
	province, err := services.GetSpecificMunicipality(provinceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, province)
}

func GetSpecificMunicipality(c *gin.Context) {
	municipalityID, err := strconv.Atoi(c.Param("municipality_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid municipality ID"})
		return
	}
	municipality, err := services.GetSpecificMunicipality(municipalityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, municipality)
}

func GetSpecificBarangay(c *gin.Context) {
	municipalityID, err := strconv.Atoi(c.Param("municipality_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid barangay ID"})
		return
	}
	barangay, err := services.GetSpecificBarangay(municipalityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barangay)
}
