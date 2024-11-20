package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func PropertyAdminRoutes(router *gin.Engine) {
	routes := router.Group("/admin/property")
	{
		routes.GET("/stands", adminService.GetAllStands)
		routes.GET("/lands", adminService.GetAllLandProperties)
		routes.GET("/commercial/rentals", adminService.GetAllCommercialRentalProperties)
		routes.GET("/commercial/onsale", adminService.GetAllCommercialForSaleProperties)
		routes.GET("/residential/onsale", adminService.GetAllResidentialForSaleProperties)
		routes.GET("/residential/rentals", adminService.GetAllResidentialRentalProperties)
		routes.GET("/media", adminService.GetAllPropertiesImagesOrVideos)
		routes.GET("/location", adminService.GetAllPropertiesLocation)
		routes.GET("/insights", adminService.GetAllPropertiesInsights)
	}
}
