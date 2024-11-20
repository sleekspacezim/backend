package manager

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
	managerRepo "SleekSpace/repositories/manager"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateManager(c *gin.Context) {
	var manager managerDtos.ManagerCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&manager)

	modelFieldsValidationError := validateModelFields.Struct(manager)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	var imageUrl string = ""
	if manager.ProfilePicture.Image != "" {
		imageUrl = <-storage.UploadFile(manager.ProfilePicture.Image, manager.ProfilePicture.Name, c)
	}
	newManager := managerModels.Manager{
		UserId: manager.UserId,
		Name:   manager.Name,
		Email:  manager.Email,
		ProfilePicture: managerModels.ManagerProfilePicture{
			Name:        manager.ProfilePicture.Name,
			Uri:         imageUrl,
			ContentType: manager.ProfilePicture.ContentType,
			FileType:    manager.ProfilePicture.FileType,
			Size:        manager.ProfilePicture.Size,
		},
		ManagerContactNumbers: []managerModels.ManagerContactNumber{
			{
				Number:       manager.Contacts[0].Number,
				CountryCode:  manager.Contacts[0].CountryCode,
				CountryAbbrv: manager.Contacts[0].CountryAbbrv,
				Type:         manager.Contacts[0].Type,
			},
			{
				Number:       manager.Contacts[1].Number,
				CountryCode:  manager.Contacts[1].CountryCode,
				CountryAbbrv: manager.Contacts[1].CountryAbbrv,
				Type:         manager.Contacts[1].Type,
			},
		},
	}
	isManagerCreated := managerRepo.CreateManager(&newManager)
	if isManagerCreated {
		createdManager := userRepo.GetUserByIdWithManager(generalUtilities.ConvertIntToString(manager.UserId))
		if createdManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.ManagerAccountCreatedButNoDataRetrievedError})
			return
		}
		createdManagerWithAssociations := managerRepo.GetManagerByManagerId(generalUtilities.ConvertIntToString(createdManager.Manager.Id))
		if createdManagerWithAssociations == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.ManagerAccountCreatedButNoDataRetrievedError})
		}
		c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(createdManagerWithAssociations)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a manager"})
	}

}

func GetManagerByUserId(c *gin.Context) {
	id := c.Param("id")
	user := userRepo.GetUserByIdWithManager(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	manager := managerRepo.GetManagerByManagerId(generalUtilities.ConvertIntToString(user.Manager.Id))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(manager)})
}

func GetManagerByManagerId(c *gin.Context) {
	id := c.Param("id")
	manager := managerRepo.GetManagerByManagerId(id)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(manager)})
}

func UpdateManagerEmailAndName(c *gin.Context) {
	managerId := c.Param("id")
	var nameAndEmailUpdates managerDtos.UpdateManagerNameAndEmailDTO
	validateModelFields := validator.New()
	c.BindJSON(&nameAndEmailUpdates)

	modelFieldsValidationError := validateModelFields.Struct(nameAndEmailUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	manager := managerRepo.GetManagerByManagerId(managerId)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}

	manager.Email = nameAndEmailUpdates.Email
	manager.Name = nameAndEmailUpdates.Name
	isManagerUpdated := managerRepo.UpdateManager(manager)
	if isManagerUpdated {
		updatedManager := managerRepo.GetManagerByManagerId(c.Param("id"))
		if updatedManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(updatedManager)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": constantsUtilities.ManagerAccountUpdateError})

}

func DeleteManager(c *gin.Context) {
	id := c.Param("id")
	manager := managerRepo.GetManagerByManagerId(id)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}
	if manager.ProfilePicture.Uri != "" {
		<-storage.DeleteFile(manager.ProfilePicture.Name, c)
	}
	isManagerDeleted := managerRepo.DeleteManagerById(manager)
	if isManagerDeleted {
		c.String(http.StatusOK, constantsUtilities.ManagerAccountDeleteSuccess)
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}
}
