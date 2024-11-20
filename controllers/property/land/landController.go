package land

import (
	"SleekSpace/middleware"
	landService "SleekSpace/services/property/land"

	"github.com/gin-gonic/gin"
)

func LandPropertyRoutes(router *gin.Engine) {
	routes := router.Group("/property/land")
	{
		routes.POST("", middleware.AuthValidator, landService.CreateLandPropertyForSale)
		routes.GET("", landService.GetAllLandPropertiesForLoggedOutUser)
		routes.GET("/:id", landService.GetLandPropertyByIdForLoggedOutUser)
		routes.GET("/logged-in", middleware.AuthValidator, landService.GetAllLandPropertiesForLoggedInUser)
		routes.GET("/logged-in/:id", middleware.AuthValidator, landService.GetLandPropertyByIdForLoggedInUser)
		routes.PUT("/:id", middleware.AuthValidator, landService.UpdateLandPropertyDetails)
		routes.DELETE("/:id", middleware.AuthValidator, landService.DeleteLandPropertyById)
	}
}
