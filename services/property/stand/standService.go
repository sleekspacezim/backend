package stand

import (
	"net/http"

	standDtos "SleekSpace/dtos/property/stand"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	userModels "SleekSpace/models/user"
	managerRepo "SleekSpace/repositories/manager"
	standRepo "SleekSpace/repositories/property/stand"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constants "SleekSpace/utilities/constants"
	favoritesUtilities "SleekSpace/utilities/funcs/favorites"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateStandForSale(c *gin.Context) {
	var standInfo standDtos.StandCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&standInfo)

	modelFieldsValidationError := validateModelFields.Struct(standInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(standInfo.Media) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(standInfo.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	mediaList := propertyUtilities.MediaListWithNoPropertyId(standInfo.Media)
	mediaUrls := storage.UploadFiles(mediaList, c)

	newStandForSale := managerModels.Stand{
		ManagerId:          standInfo.ManagerId,
		UniqueId:           propertyUtilities.GeneratePropertyUniqueId(),
		Price:              standInfo.Price,
		SizeNumber:         standInfo.SizeNumber,
		SizeDimensions:     standInfo.SizeDimensions,
		Status:             standInfo.Status,
		IsServiced:         standInfo.IsServiced,
		IsNegotiable:       standInfo.IsNegotiable,
		AreaHasElectricity: standInfo.AreaHasElectricity,
		Level:              standInfo.Level,
		Type:               standInfo.Type,
		OtherDetails:       standInfo.OtherDetails,
		Currency:           standInfo.Currency,
		MarketingStatement: standInfo.MarketingStatement,
		Manager:            *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			EmailAttempts:     0,
			CallAttempts:      0,
			WhatsAppAttempts:  0,
			PropertyType:      constants.StandPropertyType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(standInfo.Media, constants.StandPropertyType, mediaUrls),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  standInfo.PropertyLocation.Boundingbox,
			Lat:          standInfo.PropertyLocation.Lat,
			Lon:          standInfo.PropertyLocation.Lon,
			DisplayName:  standInfo.PropertyLocation.DisplayName,
			City:         standInfo.PropertyLocation.City,
			County:       standInfo.PropertyLocation.County,
			Country:      standInfo.PropertyLocation.Country,
			CountryCode:  standInfo.PropertyLocation.CountryCode,
			Province:     standInfo.PropertyLocation.Province,
			Surburb:      standInfo.PropertyLocation.Surburb,
			PropertyType: constants.StandPropertyType,
		},
	}

	isStandCreated := standRepo.CreateStandForSale(&newStandForSale)
	if isStandCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your stand was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your stand"})
	}

}

func GetAllStandsForLoggedOutUser(c *gin.Context) {
	stands := standRepo.GetAllStands(c)
	responseList := []standDtos.StandWithManagerResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandWithManagerResponse(stands[i])
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

func GetAllStandsForLoggedInUser(c *gin.Context) {
	stands := standRepo.GetAllStands(c)
	responseList := []standDtos.StandWithManagerResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandWithManagerResponse(stands[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForStandPropertyWithManager(responseList, c),
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

func UpdateStandDetails(c *gin.Context) {
	var standUpdates standDtos.StandUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&standUpdates)

	modelFieldsValidationError := validateModelFields.Struct(standUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldStandData := standRepo.GetStandById(c.Param("id"))
	if oldStandData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	oldStandData.AreaHasElectricity = standUpdates.AreaHasElectricity
	oldStandData.IsNegotiable = standUpdates.IsNegotiable
	oldStandData.Price = standUpdates.Price
	oldStandData.SizeNumber = standUpdates.SizeNumber
	oldStandData.SizeDimensions = standUpdates.SizeDimensions
	oldStandData.Level = standUpdates.Level
	oldStandData.IsServiced = standUpdates.IsServiced
	oldStandData.Status = standUpdates.Status
	oldStandData.Type = standUpdates.Type
	oldStandData.UniqueId = standUpdates.UniqueId
	oldStandData.OtherDetails = standUpdates.OtherDetails
	oldStandData.Currency = standUpdates.Currency
	oldStandData.MarketingStatement = standUpdates.MarketingStatement

	isStandUpdated := standRepo.UpdateStand(oldStandData)
	if !isStandUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "stand details update failed"})
		return
	}
	updatedStand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if updatedStand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*updatedStand)})
}

func GetStandByIdForLoggedOutUser(c *gin.Context) {
	stand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if stand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandWithManagerResponse(*stand)})
}

func GetStandByIdForLoggedInUser(c *gin.Context) {
	stand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if stand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	property := propertyUtilities.PropertyStandWithManagerResponse(*stand)
	userEmail := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(userEmail)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if len(user.FavoriteStands) > 0 {
		for i := 0; i < len(user.FavoriteStands); i++ {
			if user.FavoriteStands[i] == property.Id {
				property.IsFavorite = true
			}
		}
	} else {
		property.IsFavorite = false
	}
	c.JSON(http.StatusOK, gin.H{"response": property})
}

func GetManagerStandsByManagerId(c *gin.Context) {
	managerProperties := standRepo.GetManagerStandsByManagerId(c.Param("id"))
	propertyIdList := []int{}
	if len(managerProperties) > 0 {
		for i := 0; i < len(managerProperties); i++ {
			propertyIdList = append(propertyIdList, managerProperties[i].Id)
		}
	}
	properties := standRepo.GetStandPropertyForSaleByIds(propertyIdList, c)
	propertiesResponse := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.PropertyStandWithManagerResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": favoritesUtilities.ProcessFavoritesForStandPropertyWithManager(
			propertiesResponse, c,
		),
		"totalPages": c.GetInt("totalPages"),
		"count":      c.GetInt64("count"),
	})
}

func GetAllStandsByLocationForLoggedInUser(c *gin.Context) {
	stands := standRepo.GetAllStandsByLocation(c, c.Param("location"))
	responseList := []standDtos.StandWithManagerResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandWithManagerResponse(stands[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": favoritesUtilities.ProcessFavoritesForStandPropertyWithManager(responseList, c),
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

func GetAllStandsByLocationForLoggedOutUser(c *gin.Context) {
	stands := standRepo.GetAllStandsByLocation(c, c.Param("location"))
	responseList := []standDtos.StandWithManagerResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandWithManagerResponse(stands[i])
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

func DeleteStandById(c *gin.Context) {
	stand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if len(stand.PropertyMedia) > 0 {
		var fileNames []string
		for i := 0; i < len(stand.PropertyMedia); i++ {
			fileNames = append(fileNames, stand.PropertyMedia[i].Name)
		}
		<-storage.DeleteFiles(fileNames, c)
	}
	isStandDeleted := standRepo.DeleteStandById(c.Param("id"))
	if !isStandDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your stand was successfully deleted"})
		return
	}
}
