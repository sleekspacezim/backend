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

func GetFavoriteResidentialRentalProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialRentalProperties(user.FavoriteResidentialRentalProperties, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite residential rental properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.ResidentialRentalPropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
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

func AddFavoriteResidentialRentalProperty(c *gin.Context) {
	var residentialRentalPropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyId)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyId)
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
		user.FavoriteResidentialRentalProperties,
		residentialRentalPropertyId.Id,
	)
	user.FavoriteResidentialRentalProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialRentalProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential rental properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(processedProperties, propertyUtilities.ResidentialRentalPropertyWithManagerResponse(properties[i]))
		}
	}

	response := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
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

func RemoveFavoriteResidentialRentalProperty(c *gin.Context) {
	var residentialRentalPropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyId)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyId)
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
	for i := 0; i < len(user.FavoriteResidentialRentalProperties); i++ {
		if user.FavoriteResidentialRentalProperties[i] != residentialRentalPropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteResidentialRentalProperties[i])
		}
	}

	user.FavoriteResidentialRentalProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialRentalProperties(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential rental properties"})
		return
	}
	processedProperties := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.ResidentialRentalPropertyWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
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
