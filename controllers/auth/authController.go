package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.POST("/login", userService.Login)
		routes.GET("/logout", middleware.AuthValidator, userService.Logout)
		routes.PUT("/password", userService.UpdatePassword)
		routes.POST("/register", userService.Registration)
		routes.POST("/verification-code/registration", userService.VerifyCodeForRegistration)
		routes.POST("/verification-code/security", userService.VerifyCodeForSecurity)
		routes.POST("/verification-code", userService.CreateVerificationCode)
		routes.GET("/resend-verification-code/:id", userService.ResendVerificationCode)
	}
}
