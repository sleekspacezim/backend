package controllers

import (
	"SleekSpace/httpServices/location"

	"github.com/gin-gonic/gin"
)

func LocationIQRoutes(router *gin.Engine) {
	routes := router.Group("/locationIQ")
	{
		routes.POST("/autocomplete", location.LocationAutoComplete)
		routes.POST("/reverse-geocoding", location.LocationReverseGeoCoding)
	}
}
