package land

import (
	"net/http"

	landDtos "SleekSpace/dtos/property/land"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	userModels "SleekSpace/models/user"
	managerRepo "SleekSpace/repositories/manager"
	landRepo "SleekSpace/repositories/property/land"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constants "SleekSpace/utilities/constants"
	favoritesUtilities "SleekSpace/utilities/funcs/favorites"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateLandPropertyForSale(c *gin.Context) {
	var landDetails landDtos.LandForSalePropertyCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&landDetails)

	modelFieldsValidationError := validateModelFields.Struct(landDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(landDetails.Media) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(landDetails.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	mediaList := propertyUtilities.MediaListWithNoPropertyId(landDetails.Media)
	mediaUrls := storage.UploadFiles(mediaList, c)

	newLandForSale := managerModels.LandForSaleProperty{
		ManagerId:          landDetails.ManagerId,
		UniqueId:           propertyUtilities.GeneratePropertyUniqueId(),
		Price:              landDetails.Price,
		SizeNumber:         landDetails.SizeNumber,
		SizeDimensions:     landDetails.SizeDimensions,
		Status:             landDetails.Status,
		Type:               landDetails.Type,
		AreaHasElectricity: landDetails.AreaHasElectricity,
		HasWater:           landDetails.HasWater,
		IsNegotiable:       landDetails.IsNegotiable,
		OtherDetails:       landDetails.OtherDetails,
		Currency:           landDetails.Currency,
		MarketingStatement: landDetails.MarketingStatement,
		Manager:            *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.LandPropertyType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(landDetails.Media, constants.LandPropertyType, mediaUrls),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  landDetails.PropertyLocation.Boundingbox,
			Lat:          landDetails.PropertyLocation.Lat,
			Lon:          landDetails.PropertyLocation.Lon,
			DisplayName:  landDetails.PropertyLocation.DisplayName,
			City:         landDetails.PropertyLocation.City,
			County:       landDetails.PropertyLocation.County,
			Country:      landDetails.PropertyLocation.Country,
			CountryCode:  landDetails.PropertyLocation.CountryCode,
			Province:     landDetails.PropertyLocation.Province,
			Surburb:      landDetails.PropertyLocation.Surburb,
			PropertyType: constants.LandPropertyType,
		},
	}

	isLandCreated := landRepo.CreateLandPropertyForSale(&newLandForSale)
	if isLandCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your land was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your land"})
	}

}

func GetAllLandPropertiesForLoggedOutUser(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesForSale(c)
	responseList := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyWithManagerResponse(landProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
			"count":      c.GetInt64("count"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func GetAllLandPropertiesForLoggedInUser(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesForSale(c)
	responseList := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyWithManagerResponse(landProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForLandPropertyWithManager(responseList, c),
			"totalPages": c.GetInt("totalPages"),
			"count":      c.GetInt64("count"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func UpdateLandPropertyDetails(c *gin.Context) {
	var landUpdates landDtos.LandForSalePropertyUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&landUpdates)

	modelFieldsValidationError := validateModelFields.Struct(landUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldLandData := landRepo.GetLandPropertyForSaleById(c.Param("id"))
	if oldLandData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	oldLandData.AreaHasElectricity = landUpdates.AreaHasElectricity
	oldLandData.HasWater = landUpdates.HasWater
	oldLandData.Price = landUpdates.Price
	oldLandData.SizeNumber = landUpdates.SizeNumber
	oldLandData.SizeDimensions = landUpdates.SizeDimensions
	oldLandData.Status = landUpdates.Status
	oldLandData.Type = landUpdates.Type
	oldLandData.UniqueId = landUpdates.UniqueId
	oldLandData.IsNegotiable = landUpdates.IsNegotiable
	oldLandData.OtherDetails = landUpdates.OtherDetails
	oldLandData.Currency = landUpdates.Currency
	oldLandData.MarketingStatement = landUpdates.MarketingStatement

	isStandUpdated := landRepo.UpdateLandPropertyForSale(oldLandData)
	if !isStandUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "land details update failed"})
		return
	}
	updatedLand := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if updatedLand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyResponse(*updatedLand)})
}

func GetLandPropertyByIdForLoggedOutUser(c *gin.Context) {
	land := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if land == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyWithManagerResponse(*land)})
}

func GetLandPropertyByIdForLoggedInUser(c *gin.Context) {
	land := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if land == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	property := propertyUtilities.LandPropertyWithManagerResponse(*land)
	userEmail := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(userEmail)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if len(user.FavoriteLandProperties) > 0 {
		for i := 0; i < len(user.FavoriteLandProperties); i++ {
			if user.FavoriteLandProperties[i] == property.Id {
				property.IsFavorite = true
			}
		}
	} else {
		property.IsFavorite = false
	}
	c.JSON(http.StatusOK, gin.H{"response": property})
}

func GetManagerLandPropertiesByManagerId(c *gin.Context) {
	landProperties := landRepo.GetManagerLandPropertiesForSaleByManagerId(c.Param("id"))
	landPropertiesResponse := []landDtos.LandForSalePropertyResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			landResponse := propertyUtilities.LandPropertyResponse(landProperties[i])
			landPropertiesResponse = append(landPropertiesResponse, landResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"response": favoritesUtilities.ProcessFavoritesForLandPropertyWithoutManager(landPropertiesResponse, c),
	})
}

func GetAllLandPropertiesByLocationForLoggedInUser(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesByLocation(c, c.Param("location"))
	responseList := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyWithManagerResponse(landProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForLandPropertyWithManager(responseList, c),
			"totalPages": c.GetInt("totalPages"),
			"count":      c.GetInt64("count"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func GetAllLandPropertiesByLocationForLoggedOutUser(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesByLocation(c, c.Param("location"))
	responseList := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyWithManagerResponse(landProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
			"count":      c.GetInt64("count"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func DeleteLandPropertyById(c *gin.Context) {
	landProperty := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if len(landProperty.PropertyMedia) > 0 {
		var fileNames []string
		for i := 0; i < len(landProperty.PropertyMedia); i++ {
			fileNames = append(fileNames, landProperty.PropertyMedia[i].Name)
		}
		<-storage.DeleteFiles(fileNames, c)
	}
	isLandDeleted := landRepo.DeleteLandPropertyForSaleById(c.Param("id"))
	if !isLandDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your land was successfully deleted"})
		return
	}
}
