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

func CreateResidentialPropertyForSale(c *gin.Context) {
	var residentialPropertyForSaleDetails residentialDtos.ResidentialPropertyForSaleCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialPropertyForSaleDetails)

	modelFieldsValidationError := validateModelFields.Struct(residentialPropertyForSaleDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(residentialPropertyForSaleDetails.Media) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(residentialPropertyForSaleDetails.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	mediaList := propertyUtilities.MediaListWithNoPropertyId(residentialPropertyForSaleDetails.Media)
	mediaUrls := storage.UploadFiles(mediaList, c)

	newResidentialPropertyForSale := managerModels.ResidentialPropertyForSale{
		ManagerId:              residentialPropertyForSaleDetails.ManagerId,
		UniqueId:               propertyUtilities.GeneratePropertyUniqueId(),
		Price:                  residentialPropertyForSaleDetails.Price,
		SizeNumber:             residentialPropertyForSaleDetails.SizeNumber,
		SizeDimensions:         residentialPropertyForSaleDetails.SizeDimensions,
		Status:                 residentialPropertyForSaleDetails.Status,
		Type:                   residentialPropertyForSaleDetails.Type,
		YearBuilt:              residentialPropertyForSaleDetails.YearBuilt,
		Bedrooms:               residentialPropertyForSaleDetails.Bedrooms,
		Bathrooms:              residentialPropertyForSaleDetails.Bathrooms,
		Storeys:                residentialPropertyForSaleDetails.Storeys,
		HasElectricity:         residentialPropertyForSaleDetails.HasElectricity,
		HasWater:               residentialPropertyForSaleDetails.HasWater,
		NumberOfRooms:          residentialPropertyForSaleDetails.NumberOfRooms,
		NumberOfGarages:        residentialPropertyForSaleDetails.NumberOfGarages,
		HasSwimmingPool:        residentialPropertyForSaleDetails.HasSwimmingPool,
		IsNegotiable:           residentialPropertyForSaleDetails.IsNegotiable,
		IsPaved:                residentialPropertyForSaleDetails.IsPaved,
		IsPlustered:            residentialPropertyForSaleDetails.IsPlustered,
		IsPainted:              residentialPropertyForSaleDetails.IsPainted,
		IsTiled:                residentialPropertyForSaleDetails.IsTiled,
		HasBoreHole:            residentialPropertyForSaleDetails.HasBoreHole,
		HasCeiling:             residentialPropertyForSaleDetails.HasCeiling,
		TypeOfExteriorSecurity: residentialPropertyForSaleDetails.TypeOfExteriorSecurity,
		OtherExteriorFeatures:  residentialPropertyForSaleDetails.OtherExteriorFeatures,
		OtherInteriorFeatures:  residentialPropertyForSaleDetails.OtherInteriorFeatures,
		Currency:               residentialPropertyForSaleDetails.Currency,
		MarketingStatement:     residentialPropertyForSaleDetails.MarketingStatement,
		Manager:                *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.ResidentialPropertyForSaleType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(
			residentialPropertyForSaleDetails.Media,
			constants.ResidentialPropertyForSaleType,
			mediaUrls),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  residentialPropertyForSaleDetails.PropertyLocation.Boundingbox,
			Lat:          residentialPropertyForSaleDetails.PropertyLocation.Lat,
			Lon:          residentialPropertyForSaleDetails.PropertyLocation.Lon,
			DisplayName:  residentialPropertyForSaleDetails.PropertyLocation.DisplayName,
			City:         residentialPropertyForSaleDetails.PropertyLocation.City,
			County:       residentialPropertyForSaleDetails.PropertyLocation.County,
			Country:      residentialPropertyForSaleDetails.PropertyLocation.Country,
			CountryCode:  residentialPropertyForSaleDetails.PropertyLocation.CountryCode,
			Province:     residentialPropertyForSaleDetails.PropertyLocation.Province,
			Surburb:      residentialPropertyForSaleDetails.PropertyLocation.Surburb,
			PropertyType: constants.ResidentialPropertyForSaleType,
		},
	}

	isResidentialPropertyForSaleCreated := residentialRepo.CreateResidentialPropertyForSale(&newResidentialPropertyForSale)
	if isResidentialPropertyForSaleCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateResidentialPropertyForSaleDetails(c *gin.Context) {
	var residentialPropertyForSaleUpdates residentialDtos.ResidentialPropertyForSaleUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialPropertyForSaleUpdates)

	modelFieldsValidationError := validateModelFields.Struct(residentialPropertyForSaleUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldResidentialPropertyForSaleData := residentialRepo.GetResidentialPropertyForSaleById(c.Param("id"))
	if oldResidentialPropertyForSaleData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	oldResidentialPropertyForSaleData.IsNegotiable = residentialPropertyForSaleUpdates.IsNegotiable
	oldResidentialPropertyForSaleData.Price = residentialPropertyForSaleUpdates.Price
	oldResidentialPropertyForSaleData.SizeNumber = residentialPropertyForSaleUpdates.SizeNumber
	oldResidentialPropertyForSaleData.SizeDimensions = residentialPropertyForSaleUpdates.SizeDimensions
	oldResidentialPropertyForSaleData.Status = residentialPropertyForSaleUpdates.Status
	oldResidentialPropertyForSaleData.Type = residentialPropertyForSaleUpdates.Type
	oldResidentialPropertyForSaleData.Bathrooms = residentialPropertyForSaleUpdates.Bathrooms
	oldResidentialPropertyForSaleData.Bedrooms = residentialPropertyForSaleUpdates.Bedrooms
	oldResidentialPropertyForSaleData.HasSwimmingPool = residentialPropertyForSaleUpdates.HasSwimmingPool
	oldResidentialPropertyForSaleData.HasWater = residentialPropertyForSaleUpdates.HasWater
	oldResidentialPropertyForSaleData.HasElectricity = residentialPropertyForSaleUpdates.HasElectricity
	oldResidentialPropertyForSaleData.NumberOfRooms = residentialPropertyForSaleUpdates.NumberOfRooms
	oldResidentialPropertyForSaleData.NumberOfGarages = residentialPropertyForSaleUpdates.NumberOfGarages
	oldResidentialPropertyForSaleData.Storeys = residentialPropertyForSaleUpdates.Storeys
	oldResidentialPropertyForSaleData.YearBuilt = residentialPropertyForSaleUpdates.YearBuilt
	oldResidentialPropertyForSaleData.UniqueId = residentialPropertyForSaleUpdates.UniqueId
	oldResidentialPropertyForSaleData.OtherInteriorFeatures = residentialPropertyForSaleUpdates.OtherInteriorFeatures
	oldResidentialPropertyForSaleData.OtherExteriorFeatures = residentialPropertyForSaleUpdates.OtherExteriorFeatures
	oldResidentialPropertyForSaleData.Currency = residentialPropertyForSaleUpdates.Currency
	oldResidentialPropertyForSaleData.MarketingStatement = residentialPropertyForSaleUpdates.MarketingStatement
	oldResidentialPropertyForSaleData.IsPaved = residentialPropertyForSaleUpdates.IsPaved
	oldResidentialPropertyForSaleData.IsPlustered = residentialPropertyForSaleUpdates.IsPlustered
	oldResidentialPropertyForSaleData.IsPainted = residentialPropertyForSaleUpdates.IsPainted
	oldResidentialPropertyForSaleData.IsTiled = residentialPropertyForSaleUpdates.IsTiled
	oldResidentialPropertyForSaleData.HasBoreHole = residentialPropertyForSaleUpdates.HasBoreHole
	oldResidentialPropertyForSaleData.HasCeiling = residentialPropertyForSaleUpdates.HasCeiling
	oldResidentialPropertyForSaleData.TypeOfExteriorSecurity = residentialPropertyForSaleUpdates.TypeOfExteriorSecurity

	isResidentialPropertyForSaleUpdated := residentialRepo.UpdateResidentialPropertyForSale(oldResidentialPropertyForSaleData)
	if !isResidentialPropertyForSaleUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdatedResidentialPropertyForSale := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if UpdatedResidentialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialForSalePropertyResponse(*UpdatedResidentialPropertyForSale)})
}

func GetAllResidentialForSalePropertiesForLoggedOutUser(c *gin.Context) {
	residentialPropertiesForSale := residentialRepo.GetAllResidentialPropertiesForSale(c)
	responseList := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
	if len(residentialPropertiesForSale) > 0 {
		for i := 0; i < len(residentialPropertiesForSale); i++ {
			responseItem := propertyUtilities.ResidentialForSalePropertyWithManagerResponse(residentialPropertiesForSale[i])
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

func GetAllResidentialForSalePropertiesForLoggedInUser(c *gin.Context) {
	residentialPropertiesForSale := residentialRepo.GetAllResidentialPropertiesForSale(c)
	responseList := []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto{}
	if len(residentialPropertiesForSale) > 0 {
		for i := 0; i < len(residentialPropertiesForSale); i++ {
			responseItem := propertyUtilities.ResidentialForSalePropertyWithManagerResponse(residentialPropertiesForSale[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForResidentialForSalePropertyWithManager(responseList, c),
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

func GetResidentialPropertyForSaleByIdLoggedOutUser(c *gin.Context) {
	residentialPropertyForSale := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if residentialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": propertyUtilities.ResidentialForSalePropertyWithManagerResponse(*residentialPropertyForSale),
	})
}

func GetResidentialPropertyForSaleByIdLoggedInUser(c *gin.Context) {
	residentialPropertyForSale := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if residentialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	property := propertyUtilities.ResidentialForSalePropertyWithManagerResponse(*residentialPropertyForSale)
	userEmail := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(userEmail)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if len(user.FavoriteResidentialForSaleProperties) > 0 {
		for i := 0; i < len(user.FavoriteResidentialForSaleProperties); i++ {
			if user.FavoriteResidentialForSaleProperties[i] == property.Id {
				property.IsFavorite = true
			}
		}
	} else {
		property.IsFavorite = false
	}
	c.JSON(http.StatusOK, gin.H{"response": property})
}

func GetManagerResidentialPropertiesForSaleByManagerId(c *gin.Context) {
	properties := residentialRepo.GetManagerResidentialPropertiesForSaleByManagerId(c.Param("id"))
	propertiesResponse := []residentialDtos.ResidentialPropertyForSaleResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.ResidentialForSalePropertyResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"response": favoritesUtilities.ProcessFavoritesForResidentialForSalePropertyWithoutManager(
			propertiesResponse, c,
		),
	})
}

func DeleteResidentialPropertyForSaleById(c *gin.Context) {
	residentialPropertiesForSale := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if len(residentialPropertiesForSale.PropertyMedia) > 0 {
		var fileNames []string
		for i := 0; i < len(residentialPropertiesForSale.PropertyMedia); i++ {
			fileNames = append(fileNames, residentialPropertiesForSale.PropertyMedia[i].Name)
		}
		<-storage.DeleteFiles(fileNames, c)
	}
	isPropertyDeleted := residentialRepo.DeleteResidentialPropertyForSaleById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
