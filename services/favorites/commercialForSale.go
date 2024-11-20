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

func GetFavoriteCommercialForSaleProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialPropertiesForSale(user.FavoriteCommercialForSaleProperties, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial for sale properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.CommercialPropertyForSaleWithManagerResponse(
					properties[i],
				),
			)
		}
	}
	response := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
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

func AddFavoriteCommercialForSaleProperty(c *gin.Context) {
	var commercialForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&commercialForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(commercialForSalePropertyId)
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
		user.FavoriteCommercialForSaleProperties,
		commercialForSalePropertyId.Id,
	)
	user.FavoriteCommercialForSaleProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialPropertiesForSale(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial for sale properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.CommercialPropertyForSaleWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
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

func RemoveFavoriteCommercialForSaleProperty(c *gin.Context) {
	var commercialForSalePropertyId favoritesDtos.FavouritePropertyId
	validateModelFields := validator.New()
	c.BindJSON(&commercialForSalePropertyId)

	modelFieldsValidationError := validateModelFields.Struct(commercialForSalePropertyId)
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
	for i := 0; i < len(user.FavoriteCommercialForSaleProperties); i++ {
		if user.FavoriteCommercialForSaleProperties[i] != commercialForSalePropertyId.Id {
			newFavoritesList = append(newFavoritesList, user.FavoriteCommercialForSaleProperties[i])
		}
	}

	user.FavoriteCommercialForSaleProperties = newFavoritesList
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialPropertiesForSale(newFavoritesList, c)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have favorite commercial for sale properties"})
		return
	}
	processedProperties := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			processedProperties = append(
				processedProperties,
				propertyUtilities.CommercialPropertyForSaleWithManagerResponse(
					properties[i],
				),
			)
		}
	}

	response := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
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
