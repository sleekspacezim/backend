package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	standDtos "SleekSpace/dtos/property/stand"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteStandProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteStandProperties(user.FavoriteStands, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have stand property favorites"})
		return
	}
	processedProperties := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.PropertyStandWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []standDtos.StandWithManagerResponseDTO{}
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

func AddFavoriteStandForSaleProperty(c *gin.Context) {
	var standForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&standForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(standForSalePropertyId)
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
		user.FavoriteStands,
		standForSalePropertyId.Id,
	)
	user.FavoriteStands = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteStandProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite stands"})
		return
	}
	processedProperties := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.PropertyStandWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []standDtos.StandWithManagerResponseDTO{}
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

func RemoveFavoriteStandForSaleProperty(c *gin.Context) {
	var standForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&standForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(standForSalePropertyId)
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
	for i := 0; i < len(user.FavoriteStands); i++ {
		if user.FavoriteStands[i] != standForSalePropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteStands[i])
		}
	}

	user.FavoriteStands = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteStandProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite stands"})
		return
	}
	processedProperties := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.PropertyStandWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []standDtos.StandWithManagerResponseDTO{}
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
