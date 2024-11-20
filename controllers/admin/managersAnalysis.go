package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func ManagerAdminRoutes(router *gin.Engine) {
	routes := router.Group("/admin/managers")
	{
		routes.GET("/profile-pictures", adminService.GetAllManagersProfilePictures)
		routes.GET("/contacts", adminService.GetAllManagersContacts)
		routes.GET("", adminService.GetAllManagers)
	}
}
