package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	fmt.Print("\n Routes")

	api := router.Group("/api/v2")
	{
		api.GET("/location/region", controllers.GetAllRegions)
		api.GET("/location/region/:region_id", controllers.GetSpecificRegion)
		api.GET("/location/province", controllers.GetAllProvinces)
		api.GET("/location/province/:province_id", controllers.GetSpecificProvince)
		api.GET("/location/municipality/:municipality_id", controllers.GetSpecificMunicipality)
		api.GET("/location/barangay/:municipality_id", controllers.GetSpecificBarangay)
	}
}
