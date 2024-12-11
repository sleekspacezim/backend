package commercial

import (
	"SleekSpace/middleware"
	commercialService "SleekSpace/services/property/commercial"

	"github.com/gin-gonic/gin"
)

func CommercialRentalPropertyForSaleRoutes(router *gin.Engine) {
	routes := router.Group("/property/commercial/rentals")
	{
		routes.POST(
			"",
			middleware.AuthValidator,
			commercialService.CreateCommercialRentalProperty,
		)
		routes.GET(
			"/search/:location",
			commercialService.GetAllCommercialForSalePropertiesByLocationForLoggedOutUser,
		)
		routes.GET(
			"/search/logged-in/:location",
			middleware.AuthValidator,
			commercialService.GetAllCommercialForSalePropertiesByLocationForLoggedInUser,
		)
		routes.GET("", commercialService.GetAllCommercialRentalPropertiesForLoggedOutUser)
		routes.GET("/:id", commercialService.GetCommercialRentalPropertyIdForLoggedOutUser)
		routes.GET(
			"/logged-in",
			middleware.AuthValidator,
			commercialService.GetAllCommercialRentalPropertiesForLoggedInUser,
		)
		routes.GET(
			"/logged-in/:id",
			middleware.AuthValidator,
			commercialService.GetCommercialRentalPropertyIdForLoggedInUser,
		)
		routes.PUT(
			"/:id",
			middleware.AuthValidator,
			commercialService.UpdateCommercialRentalPropertyDetails,
		)
		routes.DELETE(
			"/:id",
			middleware.AuthValidator,
			commercialService.DeleteCommercialRentalPropertyById,
		)
	}
}
