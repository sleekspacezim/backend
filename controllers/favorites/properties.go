package favorites

import (
	"SleekSpace/middleware"
	favoritesService "SleekSpace/services/favorites"

	"github.com/gin-gonic/gin"
)

func PropertyFavoritesRoutes(router *gin.Engine) {
	routes := router.Group("/favorites/property")
	{
		routes.GET("/land/:id", middleware.AuthValidator, favoritesService.GetFavoriteLandProperties)
		routes.PUT("/land/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteLandForSaleProperty)
		routes.PUT("/land/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteLandForSaleProperty)

		routes.GET("/stand/:id", middleware.AuthValidator, favoritesService.GetFavoriteStandProperties)
		routes.PUT("/stand/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteStandForSaleProperty)
		routes.PUT("/stand/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteStandForSaleProperty)

		routes.GET("/commercial/onsale/:id", middleware.AuthValidator, favoritesService.GetFavoriteCommercialForSaleProperties)
		routes.PUT("/commercial/onsale/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteCommercialForSaleProperty)
		routes.PUT("/commercial/onsale/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteCommercialForSaleProperty)

		routes.GET("/commercial/rentals/:id", middleware.AuthValidator, favoritesService.GetFavoriteCommercialRentalProperties)
		routes.PUT("/commercial/rentals/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteCommercialRentalProperty)
		routes.PUT("/commercial/rentals/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteCommercialRentalProperty)

		routes.GET("/residential/onsale/:id", middleware.AuthValidator, favoritesService.GetFavoriteResidentialForSaleProperties)
		routes.PUT("/residential/onsale/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteResidentialForSaleProperty)
		routes.PUT("/residential/onsale/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteResidentialForSaleProperty)

		routes.GET("/residential/rentals/:id", middleware.AuthValidator, favoritesService.GetFavoriteResidentialRentalProperties)
		routes.PUT("/residential/rentals/add/:id", middleware.AuthValidator, favoritesService.AddFavoriteResidentialRentalProperty)
		routes.PUT("/residential/rentals/remove/:id", middleware.AuthValidator, favoritesService.RemoveFavoriteResidentialRentalProperty)
	}
}
