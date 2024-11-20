package report

import (
	"SleekSpace/middleware"
	reportService "SleekSpace/services/property/report"

	"github.com/gin-gonic/gin"
)

func PropertyReportRoutes(router *gin.Engine) {
	routes := router.Group("/property/report")
	{
		routes.POST("/", middleware.AuthValidator, reportService.CreatePropertyReportById)
		routes.GET("/:id", middleware.AuthValidator, reportService.GetPropertyReportById)
		routes.GET("/property/:id", middleware.AuthValidator, reportService.GetPropertyReportsByPropertyId)
		routes.GET("/manager/:id", middleware.AuthValidator, reportService.GetPropertyReportsByManagerId)
		routes.PUT("/:id", middleware.AuthValidator, reportService.UpdatePropertyReportById)
		routes.DELETE("/:id", middleware.AuthValidator, reportService.DeletePropertyReport)
	}
}
