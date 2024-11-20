package controllers

import (
	"SleekSpace/middleware"
	managerService "SleekSpace/services/manager"
	commercialService "SleekSpace/services/property/commercial"
	landService "SleekSpace/services/property/land"
	residentialService "SleekSpace/services/property/residential"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

func ManagerRoutes(router *gin.Engine) {
	routes := router.Group("/manager")
	{
		routes.POST("", middleware.AuthValidator, managerService.CreateManager)
		routes.GET("/user/:id", middleware.AuthValidator, managerService.GetManagerByUserId)
		routes.GET("/:id", middleware.AuthValidator, managerService.GetManagerByManagerId)
		routes.GET("/stands/:id", middleware.AuthValidator, standService.GetManagerStandsByManagerId)
		routes.GET("/lands/:id", middleware.AuthValidator, landService.GetManagerLandPropertiesByManagerId)
		routes.GET("/commercial/rentals/:id", middleware.AuthValidator, commercialService.GetManagerCommercialRentalPropertiesByManagerId)
		routes.GET("/commercial/onsale/:id", middleware.AuthValidator, commercialService.GetManagerCommercialPropertiesForSaleByManagerId)
		routes.GET("/residential/rentals/:id", middleware.AuthValidator, residentialService.GetManagerResidentialRentalPropertiesByManagerId)
		routes.GET("/residential/onsale/:id", middleware.AuthValidator, residentialService.GetManagerResidentialPropertiesForSaleByManagerId)
		routes.PUT("/:id", middleware.AuthValidator, managerService.UpdateManagerEmailAndName)
		routes.DELETE("/:id", middleware.AuthValidator, managerService.DeleteManager)
		routes.PUT("/contacts/:id", middleware.AuthValidator, managerService.UpdateManagerContactNumbers)
		routes.PUT("/profile-picture/:id", middleware.AuthValidator, managerService.UpdateManagerProfilePicture)
		routes.PUT("/profile-picture/remove/:id", middleware.AuthValidator, managerService.DeleteManagerProfilePicture)
	}
}
