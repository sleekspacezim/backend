package services

import (
	"net/http"

	userDtos "SleekSpace/dtos/user"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(c *gin.Context) {
	var newPasswordData = userDtos.NewPasswordDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&newPasswordData)
	modelFieldsValidationError := validateModelFields.Struct(newPasswordData)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(generalUtilities.ConvertIntToString(newPasswordData.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	if user.IsSocialsAuthenticated {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please change socials password since you signed up with it"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPasswordData.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	updateResult := userRepo.SaveUserUpdate(user)
	if !updateResult {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password update failed please try again later"})
	}
	c.JSON(http.StatusOK, gin.H{"response": "password update successful"})
}
