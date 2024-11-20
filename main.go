package main

import (
	adminController "SleekSpace/controllers/admin"
	authController "SleekSpace/controllers/auth"
	externalApiCallsController "SleekSpace/controllers/externalApiCalls"
	favoritesController "SleekSpace/controllers/favorites"
	managerController "SleekSpace/controllers/manager"
	commercialController "SleekSpace/controllers/property/commercial"
	insightsController "SleekSpace/controllers/property/insights"
	landsController "SleekSpace/controllers/property/land"
	locationController "SleekSpace/controllers/property/location"
	mediaController "SleekSpace/controllers/property/media"
	reportController "SleekSpace/controllers/property/report"
	residentialController "SleekSpace/controllers/property/residential"
	standsController "SleekSpace/controllers/property/stand"
	userController "SleekSpace/controllers/user"
	"SleekSpace/db"
	"SleekSpace/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	storage.InitializeS3()
	db.Connect()

	userController.UserRoutes(router)
	authController.AuthRoutes(router)
	externalApiCallsController.LocationIQRoutes(router)
	managerController.ManagerRoutes(router)
	adminController.PropertyAdminRoutes(router)
	adminController.UserAdminRoutes(router)
	adminController.ManagerAdminRoutes(router)
	adminController.ReportsAdminRoutes(router)
	standsController.StandsRoutes(router)
	residentialController.ResidentialPropertyForSaleRoutes(router)
	residentialController.ResidentialRentalPropertyRoutes(router)
	commercialController.CommercialPropertyForSaleRoutes(router)
	commercialController.CommercialRentalPropertyForSaleRoutes(router)
	landsController.LandPropertyRoutes(router)
	mediaController.PropertyImagesOrVideosRoutes(router)
	locationController.LocationRoutes(router)
	insightsController.PropertyInsightsRoutes(router)
	favoritesController.PropertyFavoritesRoutes(router)
	reportController.PropertyReportRoutes(router)

	router.Run()
}
