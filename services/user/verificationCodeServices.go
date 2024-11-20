package services

import (
	"net/http"
	"strconv"
	"time"

	userDtos "SleekSpace/dtos/user"
	userRepo "SleekSpace/repositories/user"
	emailService "SleekSpace/services/email"
	"SleekSpace/tokens"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	userUtilities "SleekSpace/utilities/funcs/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func VerifyCodeForRegistration(c *gin.Context) {
	var verificationInfo = userDtos.VerificationDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&verificationInfo)
	modelFieldsValidationError := validateModelFields.Struct(verificationInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	storedVerificationCode := userRepo.GetVerificationCodeByUserId(generalUtilities.ConvertIntToString(verificationInfo.UserId))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		isUserDeleted := userRepo.DeleteUserById(strconv.Itoa(storedVerificationCode.UserId))
		if !isUserDeleted {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "this user does not exist",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this verification code has expired please signup again",
			})
			return
		}
	}
	if verificationInfo.VerificationCode != storedVerificationCode.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong verification code, please try again",
		})
		return
	}
	user := userRepo.GetUserById(strconv.Itoa(storedVerificationCode.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	user.IsActive = true
	accessToken := tokens.GenerateAccessToken(user.GivenName, user.Email, user.Id)
	if accessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate your access token",
		})
		return
	}
	user.AccessToken = accessToken
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update the accessToken",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response":   userUtilities.UserResponseMapper(user, accessToken),
		"hasPayWall": constantsUtilities.IsPaywallActive,
	})
}

func CreateVerificationCode(c *gin.Context) {
	userEmail := userDtos.CreateVerificationCodeDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&userEmail)
	modelFieldsValidationError := validateModelFields.Struct(userEmail)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserByEmail(userEmail.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	verificationCode := userRepo.GetVerificationCodeByUserId(generalUtilities.ConvertIntToString(user.Id))
	verificationCode.Code = generalUtilities.GenerateVerificationCode()
	verificationCode.ExpiryDate = time.Now().Add(time.Minute * 15)
	isVerificationCodeUpdated := userRepo.UpdateVerificationCode(&verificationCode)
	if !isVerificationCodeUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate verification code",
		})
		return
	}
	isEmailSent := emailService.SendVerificationCodeEmail(
		user.Email, user.GivenName, generalUtilities.ConvertIntToString(verificationCode.Code),
	)
	if !isEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": user.Id,
	})
}

func VerifyCodeForSecurity(c *gin.Context) {
	var verificationInfo = userDtos.VerificationDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&verificationInfo)
	modelFieldsValidationError := validateModelFields.Struct(verificationInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	storedVerificationCode := userRepo.GetVerificationCodeByUserId(generalUtilities.ConvertIntToString(verificationInfo.UserId))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "this verification code has expired please resend code again",
		})
		return
	}
	if verificationInfo.VerificationCode != storedVerificationCode.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong verification code, please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": storedVerificationCode.UserId,
	})
}

func ResendVerificationCode(c *gin.Context) {
	userId := c.Param("id")
	user := userRepo.GetUserById(userId)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	verificationCode := userRepo.GetVerificationCodeByUserId(userId)
	verificationCode.Code = generalUtilities.GenerateVerificationCode()
	verificationCode.ExpiryDate = time.Now().Add(time.Minute * 15)
	isVerificationCodeUpdated := userRepo.UpdateVerificationCode(&verificationCode)
	if !isVerificationCodeUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate another code",
		})
		return
	}
	isEmailSent := emailService.SendVerificationCodeEmail(user.Email, user.GivenName, generalUtilities.ConvertIntToString(verificationCode.Code))
	if !isEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "please check your email for verification code",
	})
}
