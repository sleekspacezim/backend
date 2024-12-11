package residential

import (
	"SleekSpace/middleware"
	residentialService "SleekSpace/services/property/residential"

	"github.com/gin-gonic/gin"
)

func ResidentialRentalPropertyRoutes(router *gin.Engine) {
	routes := router.Group("/property/residential/rentals")
	{
		routes.POST("", middleware.AuthValidator, residentialService.CreateResidentialRentalProperty)
		routes.GET("/search/:location", residentialService.GetAllResidentialRentalPropertiesByLocationForLoggedOutUser)
		routes.GET("/search/logged-in/:location", middleware.AuthValidator, residentialService.GetAllResidentialRentalPropertiesByLocationForLoggedInUser)
		routes.GET("", residentialService.GetAllResidentialRentalPropertiesForLoggedOutUser)
		routes.GET("/logged-in", middleware.AuthValidator, residentialService.GetAllResidentialRentalPropertiesForLoggedInUser)
		routes.GET("/:id", residentialService.GetResidentialRentalPropertyIdForLoggedOutUser)
		routes.GET("/logged-in/:id", middleware.AuthValidator, residentialService.GetResidentialRentalPropertyIdForLoggedInUser)
		routes.PUT("/:id", middleware.AuthValidator, residentialService.UpdateResidentialRentalPropertyDetails)
		routes.DELETE("/:id", middleware.AuthValidator, residentialService.DeleteResidentialRentalPropertyById)
	}
}
