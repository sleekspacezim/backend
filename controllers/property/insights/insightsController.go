package insights

import (
	"SleekSpace/middleware"
	insightsService "SleekSpace/services/property/insights"

	"github.com/gin-gonic/gin"
)

func PropertyInsightsRoutes(router *gin.Engine) {
	routes := router.Group("/property/insights")
	{
		routes.GET("/:id", middleware.AuthValidator, insightsService.GetPropertyInsightsById)
		routes.GET("/property/:id", middleware.AuthValidator, insightsService.GetPropertyInsightsByPropertyId)
		routes.PUT("/:id", middleware.AuthValidator, insightsService.UpdatePropertyInsights)
		routes.PUT("/property/:id", insightsService.IncreamentInsightsProperties)
	}
}
