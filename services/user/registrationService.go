package services

import (
	"net/http"

	userDtos "SleekSpace/dtos/user"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	emailService "SleekSpace/services/email"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *gin.Context) {
	var userRegistartionDTO userDtos.NativeUserRegistrationDTO
	validateModelFields := validator.New()
	c.BindJSON(&userRegistartionDTO)

	modelFieldsValidationError := validateModelFields.Struct(userRegistartionDTO)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistartionDTO.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	newUser := userModels.User{
		GivenName:              userRegistartionDTO.GivenName,
		FamilyName:             userRegistartionDTO.FamilyName,
		Password:               string(hashedPassword),
		Email:                  userRegistartionDTO.Email,
		Role:                   "user",
		IsSocialsAuthenticated: false,
		RegistrationCode: userModels.VerificationCode{
			Code:       generalUtilities.GenerateVerificationCode(),
			ExpiryDate: generalUtilities.GenerateVerificationGracePeriod(),
		},
	}

	isUserCreated := userRepo.CreateUser(&newUser)
	if !isUserCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "this email already exists"})
		return
	}

	isVerificationEmailSent := emailService.SendVerificationCodeEmail(userRegistartionDTO.Email, userRegistartionDTO.GivenName, generalUtilities.ConvertIntToString(newUser.RegistrationCode.Code))
	if !isVerificationEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	createdUser := userRepo.GetUserByEmail(userRegistartionDTO.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "please check your email for verification",
		"userId":  createdUser.Id,
	})

}
