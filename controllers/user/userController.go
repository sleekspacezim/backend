package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	routes := router.Group("/user")
	{
		routes.PUT("/:id", userService.UpdateUser)
		routes.GET("/:id", middleware.AuthValidator, userService.GetUser)
		routes.GET("/email", middleware.AuthValidator, userService.GetUserByEmail)
		routes.DELETE("/:id", userService.DeleteUser)
		routes.POST("/profile-picture/:id", middleware.AuthValidator, userService.CreateUserProfilePicture)
		routes.PUT("/profile-picture/:id", middleware.AuthValidator, userService.UpdateUserProfilePicture)
		routes.DELETE("/profile-picture/:id", middleware.AuthValidator, userService.DeleteUserProfilePicture)
		routes.POST("/location", middleware.AuthValidator, userService.CreateLocation)
		routes.PUT("/location", middleware.AuthValidator, userService.UpdateLocation)
		routes.POST("/contact-number", middleware.AuthValidator, userService.CreateContactNumber)
		routes.PUT("/contact-number/:id", middleware.AuthValidator, userService.UpdateContactNumbers)
	}
}
