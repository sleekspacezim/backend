package services

import (
	"net/http"

	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateLocation(c *gin.Context) {
	var location userModels.Location
	validateModelFields := validator.New()
	c.BindJSON(&location)

	modelFieldsValidationError := validateModelFields.Struct(location)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(generalUtilities.ConvertIntToString(location.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	isLocationCreated := userRepo.CreateLocation(user, &location)
	if isLocationCreated {
		c.JSON(http.StatusOK, gin.H{"response": "location created successfully"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a location"})
	}

}

func UpdateLocation(c *gin.Context) {
	var location userModels.Location
	validateModelFields := validator.New()
	c.BindJSON(&location)
	modelFieldsValidationError := validateModelFields.Struct(location)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	isLocationUpdated := userRepo.UpdateLocation(&location)
	if !isLocationUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "location update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "location was updated successfully"})
}
