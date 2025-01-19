package residential

import (
	"net/http"

	residentialDtos "SleekSpace/dtos/property/residential"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	userModels "SleekSpace/models/user"
	managerRepo "SleekSpace/repositories/manager"
	residentialRepo "SleekSpace/repositories/property/residential"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constants "SleekSpace/utilities/constants"
	favoritesUtilities "SleekSpace/utilities/funcs/favorites"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateResidentialRentalProperty(c *gin.Context) {
	var residentialRentalPropertyDetails residentialDtos.ResidentialPropertyForRentCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyDetails)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(residentialRentalPropertyDetails.Media) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(residentialRentalPropertyDetails.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	mediaList := propertyUtilities.MediaListWithNoPropertyId(residentialRentalPropertyDetails.Media)
	mediaUrls := storage.UploadFiles(mediaList, c)

	newResidentialRentalProperty := managerModels.ResidentialRentalProperty{
		ManagerId:              residentialRentalPropertyDetails.ManagerId,
		UniqueId:               propertyUtilities.GeneratePropertyUniqueId(),
		RentAmount:             residentialRentalPropertyDetails.RentAmount,
		SizeNumber:             residentialRentalPropertyDetails.SizeNumber,
		SizeDimensions:         residentialRentalPropertyDetails.SizeDimensions,
		Status:                 residentialRentalPropertyDetails.Status,
		Type:                   residentialRentalPropertyDetails.Type,
		YearBuilt:              residentialRentalPropertyDetails.YearBuilt,
		Bedrooms:               residentialRentalPropertyDetails.Bedrooms,
		Bathrooms:              residentialRentalPropertyDetails.Bathrooms,
		Storeys:                residentialRentalPropertyDetails.Storeys,
		HasElectricity:         residentialRentalPropertyDetails.HasElectricity,
		HasWater:               residentialRentalPropertyDetails.HasWater,
		NumberOfRoomsToLet:     residentialRentalPropertyDetails.NumberOfRoomsToLet,
		NumberOfGarages:        residentialRentalPropertyDetails.NumberOfGarages,
		HasSwimmingPool:        residentialRentalPropertyDetails.HasSwimmingPool,
		IsFullHouse:            residentialRentalPropertyDetails.IsFullHouse,
		OtherExteriorFeatures:  residentialRentalPropertyDetails.OtherExteriorFeatures,
		OtherInteriorFeatures:  residentialRentalPropertyDetails.OtherInteriorFeatures,
		Currency:               residentialRentalPropertyDetails.Currency,
		MarketingStatement:     residentialRentalPropertyDetails.MarketingStatement,
		IsPaved:                residentialRentalPropertyDetails.IsPaved,
		IsPlustered:            residentialRentalPropertyDetails.IsPlustered,
		IsPainted:              residentialRentalPropertyDetails.IsPainted,
		IsTiled:                residentialRentalPropertyDetails.IsTiled,
		HasBoreHole:            residentialRentalPropertyDetails.HasBoreHole,
		HasCeiling:             residentialRentalPropertyDetails.HasCeiling,
		NumberOfRooms:          residentialRentalPropertyDetails.NumberOfRooms,
		TypeOfExteriorSecurity: residentialRentalPropertyDetails.TypeOfExteriorSecurity,
		TenantRequirements:     residentialRentalPropertyDetails.TenantRequirements,
		Manager:                *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.ResidentialRentalPropertyType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(residentialRentalPropertyDetails.Media, constants.ResidentialRentalPropertyType, mediaUrls),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  residentialRentalPropertyDetails.PropertyLocation.Boundingbox,
			Lat:          residentialRentalPropertyDetails.PropertyLocation.Lat,
			Lon:          residentialRentalPropertyDetails.PropertyLocation.Lon,
			DisplayName:  residentialRentalPropertyDetails.PropertyLocation.DisplayName,
			City:         residentialRentalPropertyDetails.PropertyLocation.City,
			County:       residentialRentalPropertyDetails.PropertyLocation.County,
			Country:      residentialRentalPropertyDetails.PropertyLocation.Country,
			CountryCode:  residentialRentalPropertyDetails.PropertyLocation.CountryCode,
			Province:     residentialRentalPropertyDetails.PropertyLocation.Province,
			Surburb:      residentialRentalPropertyDetails.PropertyLocation.Surburb,
			PropertyType: constants.ResidentialRentalPropertyType,
		},
	}

	isResidentialRentalPropertyCreated := residentialRepo.CreateResidentialRentalProperty(&newResidentialRentalProperty)
	if isResidentialRentalPropertyCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateResidentialRentalPropertyDetails(c *gin.Context) {
	var residentialRentalPropertyUpdates residentialDtos.ResidentialPropertyForRentUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyUpdates)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldResidentialRentalPropertyData := residentialRepo.GetResidentialRentalPropertyById(c.Param("id"))
	if oldResidentialRentalPropertyData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}

	oldResidentialRentalPropertyData.IsFullHouse = residentialRentalPropertyUpdates.IsFullHouse
	oldResidentialRentalPropertyData.RentAmount = residentialRentalPropertyUpdates.RentAmount
	oldResidentialRentalPropertyData.SizeNumber = residentialRentalPropertyUpdates.SizeNumber
	oldResidentialRentalPropertyData.SizeDimensions = residentialRentalPropertyUpdates.SizeDimensions
	oldResidentialRentalPropertyData.Status = residentialRentalPropertyUpdates.Status
	oldResidentialRentalPropertyData.Type = residentialRentalPropertyUpdates.Type
	oldResidentialRentalPropertyData.Bathrooms = residentialRentalPropertyUpdates.Bathrooms
	oldResidentialRentalPropertyData.Bedrooms = residentialRentalPropertyUpdates.Bedrooms
	oldResidentialRentalPropertyData.HasSwimmingPool = residentialRentalPropertyUpdates.HasSwimmingPool
	oldResidentialRentalPropertyData.HasWater = residentialRentalPropertyUpdates.HasWater
	oldResidentialRentalPropertyData.HasElectricity = residentialRentalPropertyUpdates.HasElectricity
	oldResidentialRentalPropertyData.NumberOfRoomsToLet = residentialRentalPropertyUpdates.NumberOfRoomsToLet
	oldResidentialRentalPropertyData.NumberOfGarages = residentialRentalPropertyUpdates.NumberOfGarages
	oldResidentialRentalPropertyData.Storeys = residentialRentalPropertyUpdates.Storeys
	oldResidentialRentalPropertyData.YearBuilt = residentialRentalPropertyUpdates.YearBuilt
	oldResidentialRentalPropertyData.UniqueId = residentialRentalPropertyUpdates.UniqueId
	oldResidentialRentalPropertyData.OtherInteriorFeatures = residentialRentalPropertyUpdates.OtherInteriorFeatures
	oldResidentialRentalPropertyData.OtherExteriorFeatures = residentialRentalPropertyUpdates.OtherExteriorFeatures
	oldResidentialRentalPropertyData.IsPaved = residentialRentalPropertyUpdates.IsPaved
	oldResidentialRentalPropertyData.IsPlustered = residentialRentalPropertyUpdates.IsPlustered
	oldResidentialRentalPropertyData.IsPainted = residentialRentalPropertyUpdates.IsPainted
	oldResidentialRentalPropertyData.IsTiled = residentialRentalPropertyUpdates.IsTiled
	oldResidentialRentalPropertyData.HasBoreHole = residentialRentalPropertyUpdates.HasBoreHole
	oldResidentialRentalPropertyData.HasCeiling = residentialRentalPropertyUpdates.HasCeiling
	oldResidentialRentalPropertyData.NumberOfRooms = residentialRentalPropertyUpdates.NumberOfRooms
	oldResidentialRentalPropertyData.TypeOfExteriorSecurity = residentialRentalPropertyUpdates.TypeOfExteriorSecurity
	oldResidentialRentalPropertyData.TenantRequirements = residentialRentalPropertyUpdates.TenantRequirements
	oldResidentialRentalPropertyData.Currency = residentialRentalPropertyUpdates.Currency
	oldResidentialRentalPropertyData.MarketingStatement = residentialRentalPropertyUpdates.MarketingStatement

	isResidentialRentalPropertyUpdated := residentialRepo.UpdateResidentialRentalProperty(oldResidentialRentalPropertyData)
	if !isResidentialRentalPropertyUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdateResidentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if UpdateResidentialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialRentalPropertyResponse(*UpdateResidentialRentalProperty)})
}

func GetAllResidentialRentalPropertiesForLoggedInUser(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalProperties(c)
	responseList := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(residentialRentalProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForResidentialRentalPropertyWithManager(responseList, c),
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

func GetAllResidentialRentalPropertiesForLoggedOutUser(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalProperties(c)
	responseList := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(residentialRentalProperties[i])
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

func GetAllResidentialRentalPropertiesByLocationForLoggedInUser(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalPropertiesByLocation(c, c.Param("location"))
	responseList := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(residentialRentalProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForResidentialRentalPropertyWithManager(responseList, c),
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

func GetAllResidentialRentalPropertiesByLocationForLoggedOutUser(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalPropertiesByLocation(c, c.Param("location"))
	responseList := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(residentialRentalProperties[i])
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

func GetResidentialRentalPropertyIdForLoggedOutUser(c *gin.Context) {
	residentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if residentialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": propertyUtilities.ResidentialRentalPropertyWithManagerResponse(*residentialRentalProperty),
	})
}

func GetResidentialRentalPropertyIdForLoggedInUser(c *gin.Context) {
	residentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if residentialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	property := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(*residentialRentalProperty)
	userEmail := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(userEmail)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if len(user.FavoriteResidentialRentalProperties) > 0 {
		for i := 0; i < len(user.FavoriteResidentialRentalProperties); i++ {
			if user.FavoriteResidentialRentalProperties[i] == property.Id {
				property.IsFavorite = true
			}
		}
	} else {
		property.IsFavorite = false
	}
	c.JSON(http.StatusOK, gin.H{"response": property})
}

func GetManagerResidentialRentalPropertiesByManagerId(c *gin.Context) {
	managerProperties := residentialRepo.GetManagerResidentialRentalPropertiesByManagerId(c.Param("id"))
	propertyIdList := []int{}
	if len(managerProperties) > 0 {
		for i := 0; i < len(managerProperties); i++ {
			propertyIdList = append(propertyIdList, managerProperties[i].Id)
		}
	}
	properties := residentialRepo.GetResidentialRentalPropertyByIds(propertyIdList, c)
	propertiesResponse := []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.ResidentialRentalPropertyWithManagerResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": favoritesUtilities.ProcessFavoritesForResidentialRentalPropertyWithManager(
			propertiesResponse, c,
		),
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func DeleteResidentialRentalPropertyById(c *gin.Context) {
	residentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if len(residentialRentalProperty.PropertyMedia) > 0 {
		var fileNames []string
		for i := 0; i < len(residentialRentalProperty.PropertyMedia); i++ {
			fileNames = append(fileNames, residentialRentalProperty.PropertyMedia[i].Name)
		}
		<-storage.DeleteFiles(fileNames, c)
	}
	isPropertyDeleted := residentialRepo.DeleteResidentialRentalPropertyById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
