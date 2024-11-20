package media

import (
	"SleekSpace/middleware"
	mediaService "SleekSpace/services/property/media"

	"github.com/gin-gonic/gin"
)

func PropertyImagesOrVideosRoutes(router *gin.Engine) {
	routes := router.Group("/property/media")
	{
		routes.POST("", middleware.AuthValidator, mediaService.CreatePropertyImageOrVideoWithPropertyId)
		routes.GET("/:id", middleware.AuthValidator, mediaService.GetPropertyImageOrVideoById)
		routes.PUT("/:id", middleware.AuthValidator, mediaService.UpdatePropertyImageOrVideo)
		routes.DELETE("/:id", middleware.AuthValidator, mediaService.DeletePropertyImageOrVideo)
	}
}
