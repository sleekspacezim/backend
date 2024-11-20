package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	landDtos "SleekSpace/dtos/property/land"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteLandProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteLandProperties(user.FavoriteLandProperties, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have land property favorites"})
		return
	}
	processedProperties := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.LandPropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(processedProperties) > 0 {
		for i := 0; i < len(processedProperties); i++ {
			processedProperties[i].IsFavorite = true
			response = append(response, processedProperties[i])
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"properties": response,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func AddFavoriteLandForSaleProperty(c *gin.Context) {
	var landForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&landForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(landForSalePropertyId)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	newFavoritesList := append(
		user.FavoriteLandProperties,
		landForSalePropertyId.Id,
	)
	user.FavoriteLandProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteLandProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite land"})
		return
	}
	processedProperties := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.LandPropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(processedProperties) > 0 {
		for i := 0; i < len(processedProperties); i++ {
			processedProperties[i].IsFavorite = true
			response = append(response, processedProperties[i])
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"properties": response,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func RemoveFavoriteLandForSaleProperty(c *gin.Context) {
	var landForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&landForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(landForSalePropertyId)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	newFavoritesList := []int{}
	for i := 0; i < len(user.FavoriteLandProperties); i++ {
		if user.FavoriteLandProperties[i] != landForSalePropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteLandProperties[i])
		}
	}

	user.FavoriteLandProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteLandProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite lands"})
		return
	}
	processedProperties := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.LandPropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(processedProperties) > 0 {
		for i := 0; i < len(processedProperties); i++ {
			processedProperties[i].IsFavorite = true
			response = append(response, processedProperties[i])
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"properties": response,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}
