package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	commercialDtos "SleekSpace/dtos/property/commercial"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteCommercialRentalProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialRentalProperties(user.FavoriteCommercialRentalProperties, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial rental properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.CommercialPropertyForRentWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
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

func AddFavoriteCommercialRentalProperty(c *gin.Context) {
	var commercialRentalPropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&commercialRentalPropertyId)

	modelFieldsValidationError := validateModelFields.Struct(commercialRentalPropertyId)
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
		user.FavoriteCommercialRentalProperties,
		commercialRentalPropertyId.Id,
	)
	user.FavoriteCommercialRentalProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialRentalProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial rental properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(processedProperties, propertyUtilities.CommercialPropertyForRentWithManagerResponse(properties[i]))
		}
	}

	response := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
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

func RemoveFavoriteCommercialRentalProperty(c *gin.Context) {
	var commercialRentalPropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&commercialRentalPropertyId)

	modelFieldsValidationError := validateModelFields.Struct(commercialRentalPropertyId)
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
	for i := 0; i < len(user.FavoriteCommercialRentalProperties); i++ {
		if user.FavoriteCommercialRentalProperties[i] != commercialRentalPropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteCommercialRentalProperties[i])
		}
	}

	user.FavoriteCommercialRentalProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialRentalProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial rental properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.CommercialPropertyForRentWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []commercialDtos.CommercialForRentPropertyWithManagerResponseDto{}
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
