package commercial

import (
	"net/http"

	commercialDtos "SleekSpace/dtos/property/commercial"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	userModels "SleekSpace/models/user"
	managerRepo "SleekSpace/repositories/manager"
	commercialRepo "SleekSpace/repositories/property/commercial"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constants "SleekSpace/utilities/constants"
	favoritesUtilities "SleekSpace/utilities/funcs/favorites"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateCommercialPropertyForSale(c *gin.Context) {
	var commercialPropertyForSaleDetails commercialDtos.CommercialForSalePropertyCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialPropertyForSaleDetails)

	modelFieldsValidationError := validateModelFields.Struct(commercialPropertyForSaleDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(commercialPropertyForSaleDetails.Media) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(commercialPropertyForSaleDetails.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	mediaList := propertyUtilities.MediaListWithNoPropertyId(commercialPropertyForSaleDetails.Media)
	mediaUrls := storage.UploadFiles(mediaList, c)

	newCommercialPropertyForSale := managerModels.CommercialForSaleProperty{
		ManagerId:             commercialPropertyForSaleDetails.ManagerId,
		UniqueId:              propertyUtilities.GeneratePropertyUniqueId(),
		Price:                 commercialPropertyForSaleDetails.Price,
		SizeNumber:            commercialPropertyForSaleDetails.SizeNumber,
		SizeDimensions:        commercialPropertyForSaleDetails.SizeDimensions,
		Status:                commercialPropertyForSaleDetails.Status,
		Type:                  commercialPropertyForSaleDetails.Type,
		YearBuilt:             commercialPropertyForSaleDetails.YearBuilt,
		Storeys:               commercialPropertyForSaleDetails.Storeys,
		HasElectricity:        commercialPropertyForSaleDetails.HasElectricity,
		HasWater:              commercialPropertyForSaleDetails.HasWater,
		IsNegotiable:          commercialPropertyForSaleDetails.IsNegotiable,
		NumberOfRooms:         commercialPropertyForSaleDetails.NumberOfRooms,
		OtherExteriorFeatures: commercialPropertyForSaleDetails.OtherExteriorFeatures,
		OtherInteriorFeatures: commercialPropertyForSaleDetails.OtherInteriorFeatures,
		Currency:              commercialPropertyForSaleDetails.Currency,
		MarketingStatement:    commercialPropertyForSaleDetails.MarketingStatement,
		Manager:               *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			EmailAttempts:     0,
			CallAttempts:      0,
			WhatsAppAttempts:  0,
			PropertyType:      constants.CommercialPropertyForSaleType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(commercialPropertyForSaleDetails.Media, constants.CommercialPropertyForSaleType, mediaUrls),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  commercialPropertyForSaleDetails.PropertyLocation.Boundingbox,
			Lat:          commercialPropertyForSaleDetails.PropertyLocation.Lat,
			Lon:          commercialPropertyForSaleDetails.PropertyLocation.Lon,
			DisplayName:  commercialPropertyForSaleDetails.PropertyLocation.DisplayName,
			City:         commercialPropertyForSaleDetails.PropertyLocation.City,
			County:       commercialPropertyForSaleDetails.PropertyLocation.County,
			Country:      commercialPropertyForSaleDetails.PropertyLocation.Country,
			CountryCode:  commercialPropertyForSaleDetails.PropertyLocation.CountryCode,
			Province:     commercialPropertyForSaleDetails.PropertyLocation.Province,
			Surburb:      commercialPropertyForSaleDetails.PropertyLocation.Surburb,
			PropertyType: constants.CommercialPropertyForSaleType,
		},
	}

	isCommercialPropertyForSaleCreated := commercialRepo.CreateCommercialPropertyForSale(&newCommercialPropertyForSale)
	if isCommercialPropertyForSaleCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateCommercialPropertyForSaleDetails(c *gin.Context) {
	var commercialPropertyForSaleUpdates commercialDtos.CommercialForSalePropertyUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialPropertyForSaleUpdates)

	modelFieldsValidationError := validateModelFields.Struct(commercialPropertyForSaleUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldCommercialPropertyForSaleData := commercialRepo.GetCommercialPropertyForSaleById(c.Param("id"))
	if oldCommercialPropertyForSaleData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}

	oldCommercialPropertyForSaleData.Price = commercialPropertyForSaleUpdates.Price
	oldCommercialPropertyForSaleData.SizeNumber = commercialPropertyForSaleUpdates.SizeNumber
	oldCommercialPropertyForSaleData.SizeDimensions = commercialPropertyForSaleUpdates.SizeDimensions
	oldCommercialPropertyForSaleData.Status = commercialPropertyForSaleUpdates.Status
	oldCommercialPropertyForSaleData.Type = commercialPropertyForSaleUpdates.Type
	oldCommercialPropertyForSaleData.HasWater = commercialPropertyForSaleUpdates.HasWater
	oldCommercialPropertyForSaleData.HasElectricity = commercialPropertyForSaleUpdates.HasElectricity
	oldCommercialPropertyForSaleData.NumberOfRooms = commercialPropertyForSaleUpdates.NumberOfRooms
	oldCommercialPropertyForSaleData.Storeys = commercialPropertyForSaleUpdates.Storeys
	oldCommercialPropertyForSaleData.YearBuilt = commercialPropertyForSaleUpdates.YearBuilt
	oldCommercialPropertyForSaleData.UniqueId = commercialPropertyForSaleUpdates.UniqueId
	oldCommercialPropertyForSaleData.IsNegotiable = commercialPropertyForSaleUpdates.IsNegotiable
	oldCommercialPropertyForSaleData.OtherInteriorFeatures = commercialPropertyForSaleUpdates.OtherInteriorFeatures
	oldCommercialPropertyForSaleData.OtherExteriorFeatures = commercialPropertyForSaleUpdates.OtherExteriorFeatures
	oldCommercialPropertyForSaleData.Currency = commercialPropertyForSaleUpdates.Currency
	oldCommercialPropertyForSaleData.MarketingStatement = commercialPropertyForSaleUpdates.MarketingStatement

	isCommercialPropertyForSaleUpdated := commercialRepo.UpdateCommercialPropertyForSale(oldCommercialPropertyForSaleData)
	if !isCommercialPropertyForSaleUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdatedCommercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if UpdatedCommercialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForSaleResponse(*UpdatedCommercialPropertyForSale)})
}

func GetAllCommercialForSalePropertiesForLoggedOutUser(c *gin.Context) {
	commercialPropertiesForSale := commercialRepo.GetAllCommercialPropertiesForSale(c)
	responseList := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(commercialPropertiesForSale) > 0 {
		for i := 0; i < len(commercialPropertiesForSale); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(commercialPropertiesForSale[i])
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

func GetAllCommercialForSalePropertiesForLoggedInUser(c *gin.Context) {
	commercialPropertiesForSale := commercialRepo.GetAllCommercialPropertiesForSale(c)
	responseList := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(commercialPropertiesForSale) > 0 {
		for i := 0; i < len(commercialPropertiesForSale); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(commercialPropertiesForSale[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForCommercialForSalePropertyWithManager(responseList, c),
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

func GetAllCommercialForSalePropertiesByLocationForLoggedInUser(c *gin.Context) {
	commercialForSaleProperties := commercialRepo.GetAllCommercialPropertiesForSaleByLocation(c, c.Param("location"))
	responseList := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(commercialForSaleProperties) > 0 {
		for i := 0; i < len(commercialForSaleProperties); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(commercialForSaleProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForCommercialForSalePropertyWithManager(responseList, c),
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

func GetAllCommercialForSalePropertiesByLocationForLoggedOutUser(c *gin.Context) {
	commercialForSaleProperties := commercialRepo.GetAllCommercialPropertiesForSaleByLocation(c, c.Param("location"))
	responseList := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(commercialForSaleProperties) > 0 {
		for i := 0; i < len(commercialForSaleProperties); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(commercialForSaleProperties[i])
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

func GetCommercialPropertyForSaleByIdForLoggedOutUser(c *gin.Context) {
	commercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if commercialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": propertyUtilities.CommercialPropertyForSaleWithManagerResponse(*commercialPropertyForSale),
	})
}

func GetCommercialPropertyForSaleByIdForLoggedInUser(c *gin.Context) {
	commercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if commercialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	property := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(*commercialPropertyForSale)
	userEmail := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(userEmail)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if len(user.FavoriteCommercialForSaleProperties) > 0 {
		for i := 0; i < len(user.FavoriteCommercialForSaleProperties); i++ {
			if user.FavoriteCommercialForSaleProperties[i] == property.Id {
				property.IsFavorite = true
			}
		}
	} else {
		property.IsFavorite = false
	}
	c.JSON(http.StatusOK, gin.H{"response": property})
}

func GetManagerCommercialPropertiesForSaleByManagerId(c *gin.Context) {
	managerProperties := commercialRepo.GetManagerCommercialPropertiesForSaleByManagerId(c.Param("id"))
	propertyIdList := []int{}
	if len(managerProperties) > 0 {
		for i := 0; i < len(managerProperties); i++ {
			propertyIdList = append(propertyIdList, managerProperties[i].Id)
		}
	}
	properties := commercialRepo.GetCommercialPropertyForSaleByIds(propertyIdList, c)
	propertiesResponse := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": favoritesUtilities.ProcessFavoritesForCommercialForSalePropertyWithManager(
			propertiesResponse, c,
		),
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func DeleteCommercialPropertyForSaleById(c *gin.Context) {
	commercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if len(commercialPropertyForSale.PropertyMedia) > 0 {
		var fileNames []string
		for i := 0; i < len(commercialPropertyForSale.PropertyMedia); i++ {
			fileNames = append(fileNames, commercialPropertyForSale.PropertyMedia[i].Name)
		}
		<-storage.DeleteFiles(fileNames, c)
	}
	isPropertyDeleted := commercialRepo.DeleteCommercialPropertyForSaleById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
