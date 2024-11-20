package stand

import (
	"SleekSpace/middleware"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

func StandsRoutes(router *gin.Engine) {
	routes := router.Group("/property/stand")
	{
		routes.POST("", middleware.AuthValidator, standService.CreateStandForSale)
		routes.GET("", standService.GetAllStandsForLoggedOutUser)
		routes.GET("/:id", standService.GetStandByIdForLoggedOutUser)
		routes.GET("/logged-in", middleware.AuthValidator, standService.GetAllStandsForLoggedInUser)
		routes.GET("/logged-in/:id", middleware.AuthValidator, standService.GetStandByIdForLoggedInUser)
		routes.PUT("/:id", middleware.AuthValidator, standService.UpdateStandDetails)
		routes.DELETE("/:id", middleware.AuthValidator, standService.DeleteStandById)
	}
}
