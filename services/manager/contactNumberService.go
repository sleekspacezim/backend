package manager

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerRepo "SleekSpace/repositories/manager"
	managerUtilities "SleekSpace/utilities/funcs/manager"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateManagerContactNumbers(c *gin.Context) {
	managerId := c.Param("id")
	var contactsUpdates managerDtos.ManagerContactNumbersCreationAndUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&contactsUpdates)

	modelFieldsValidationError := validateModelFields.Struct(contactsUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	manager := managerRepo.GetManagerByManagerId(managerId)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}

	areContactsCreated := managerRepo.UpdateManagerContactNumbers(manager, contactsUpdates.Contacts)
	if !areContactsCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property management account contacts failed to create"})
		return
	}

	updatedManager := managerRepo.GetManagerByManagerId(managerId)
	if updatedManager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(updatedManager)})
}

func UpdateManagerContactNumbers(c *gin.Context) {
	managerId := c.Param("id")
	var contactsUpdates managerDtos.ManagerContactNumbersCreationAndUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&contactsUpdates)

	modelFieldsValidationError := validateModelFields.Struct(contactsUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	manager := managerRepo.GetManagerByManagerId(managerId)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}

	hasContactsUpdated := managerRepo.UpdateManagerContactNumbers(manager, contactsUpdates.Contacts)
	if !hasContactsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property management account contacts failed to update"})
		return
	}

	updatedManager := managerRepo.GetManagerByManagerId(managerId)
	if updatedManager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(updatedManager)})
}
