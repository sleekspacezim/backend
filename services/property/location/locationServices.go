package location

import (
	"net/http"

	locationDtos "SleekSpace/dtos/property/location"
	propertyModels "SleekSpace/models/property"
	locationRepo "SleekSpace/repositories/property/location"
	generalServices "SleekSpace/services/property"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetPropertyLocationById(c *gin.Context) {
	location := locationRepo.GetPropertyLocationById(c.Param("id"))
	if location == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property location does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyLocationResponse(*location)})
}

func UpdatePropertyLocation(c *gin.Context) {
	var locationUpdateDetails locationDtos.PropertyLocationUpdateAndResponseDto
	validateModelFields := validator.New()
	c.BindJSON(&locationUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(locationUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	locationUpdate := propertyModels.PropertyLocation{
		Id:           locationUpdateDetails.Id,
		PropertyId:   locationUpdateDetails.PropertyId,
		PropertyType: locationUpdateDetails.PropertyType,
		Boundingbox:  locationUpdateDetails.Boundingbox,
		Lat:          locationUpdateDetails.Lat,
		Lon:          locationUpdateDetails.Lon,
		DisplayName:  locationUpdateDetails.DisplayName,
		City:         locationUpdateDetails.City,
		County:       locationUpdateDetails.County,
		Country:      locationUpdateDetails.Country,
		CountryCode:  locationUpdateDetails.CountryCode,
		Province:     locationUpdateDetails.Province,
		Surburb:      locationUpdateDetails.Surburb,
	}

	isLocationUpdated := locationRepo.UpdatePropertyLocation(&locationUpdate)
	if !isLocationUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property location"})
		return
	}
	generalServices.GetPropertyTypeById(c, locationUpdateDetails.PropertyType, locationUpdate.PropertyId)
}
