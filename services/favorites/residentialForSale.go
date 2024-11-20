package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	residentialDtos "SleekSpace/dtos/property/residential"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteResidentialForSaleProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialPropertiesForSale(user.FavoriteResidentialForSaleProperties, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite residential for sale properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.ResidentialForSalePropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
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

func AddFavoriteResidentialForSaleProperty(c *gin.Context) {
	var residentialForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&residentialForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(residentialForSalePropertyId)
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
		user.FavoriteResidentialForSaleProperties,
		residentialForSalePropertyId.Id,
	)
	user.FavoriteResidentialForSaleProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialPropertiesForSale(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite residential for sale properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.ResidentialForSalePropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
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

func RemoveFavoriteResidentialForSaleProperty(c *gin.Context) {
	var residentialForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&residentialForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(residentialForSalePropertyId)
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
	for i := 0; i < len(user.FavoriteResidentialForSaleProperties); i++ {
		if user.FavoriteResidentialForSaleProperties[i] != residentialForSalePropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteResidentialForSaleProperties[i])
		}
	}

	user.FavoriteResidentialForSaleProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialPropertiesForSale(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite residential for sale properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.ResidentialForSalePropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
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
