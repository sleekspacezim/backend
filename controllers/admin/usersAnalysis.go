package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func UserAdminRoutes(router *gin.Engine) {
	routes := router.Group("/admin")
	{
		routes.GET("/users", adminService.GetAllUsers)
		routes.GET("/users/location", adminService.GetAllUsersLocations)
		routes.GET("/users/profile-pictures", adminService.GetAllUsersProfilePictures)
		routes.GET("/verification-code/:id", adminService.GetVerificationCodeById)
		routes.GET("/verification-codes", adminService.GetAllVerificationCodes)
		routes.DELETE("/verification-code/:id", adminService.DeleteVerificationCode)
		routes.GET("/user/contacts", adminService.GetUserContacts)
	}
}
