package services

import (
	"net/http"

	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"github.com/gin-gonic/gin"
)

func CreateContactNumber(c *gin.Context) {
	var contact userModels.ContactNumber
	c.BindJSON(&contact)
	user := userRepo.GetUserById(generalUtilities.ConvertIntToString(contact.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	result := userRepo.CreateContactNumber(user, &contact)
	if result {
		c.String(http.StatusOK, "contact added")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "an error occured"})
	}
}

func UpdateContactNumbers(c *gin.Context) {
	type ContactNumbers struct {
		Contacts []userModels.ContactNumber `json:"contacts"`
	}
	var contacts ContactNumbers
	c.BindJSON(&contacts)
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	isContactsUpdated := userRepo.UpdateUserContactNumbers(user, contacts.Contacts)
	if !isContactsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "contacts update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "contacts were updated succesfully"})
}
