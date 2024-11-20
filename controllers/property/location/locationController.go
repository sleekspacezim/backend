package location

import (
	"SleekSpace/middleware"
	locationService "SleekSpace/services/property/location"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(router *gin.Engine) {
	routes := router.Group("/property/location")
	{
		routes.GET("/:id", middleware.AuthValidator, locationService.GetPropertyLocationById)
		routes.PUT("/:id", middleware.AuthValidator, locationService.UpdatePropertyLocation)
	}
}
