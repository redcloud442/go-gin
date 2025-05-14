package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inikhildubey/GoLang-Gin-boilerplate/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	fmt.Print("\n Routes")
	api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/DBusers", controllers.FetchUsers)
		api.POST("/DBusers", controllers.CreateUserController)
	}
}
