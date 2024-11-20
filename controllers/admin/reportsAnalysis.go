package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func ReportsAdminRoutes(router *gin.Engine) {
	routes := router.Group("/admin/reports")
	{
		routes.GET("", adminService.GetAllPropertiesReports)
	}
}
