package services

import (
	"net/http"

	userDtos "SleekSpace/dtos/user"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/storage"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	userUtilities "SleekSpace/utilities/funcs/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUserProfilePicture(c *gin.Context) {
	var userProfilePicture userDtos.UserProfilePictureCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&userProfilePicture)

	modelFieldsValidationError := validateModelFields.Struct(userProfilePicture)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(generalUtilities.ConvertIntToString(userProfilePicture.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	imageUrl := <-storage.UploadFile(userProfilePicture.Image, userProfilePicture.Name, c)
	newProfilePicture := userModels.UserProfilePicture{
		UserId:      userProfilePicture.UserId,
		Uri:         imageUrl,
		Name:        userProfilePicture.Name,
		FileType:    userProfilePicture.FileType,
		ContentType: userProfilePicture.ContentType,
		Size:        userProfilePicture.Size,
	}
	isUserProfilePictureCreated := userRepo.CreateUserProfilePicture(user, &newProfilePicture)
	if isUserProfilePictureCreated {
		updatedUser := userRepo.GetUserById(generalUtilities.ConvertIntToString(userProfilePicture.UserId))
		if updatedUser == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(updatedUser, updatedUser.AccessToken)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a profile picture"})
	}

}

func UpdateUserProfilePicture(c *gin.Context) {
	userId := c.Param("id")
	var profilePictureUpdate userDtos.UserProfilePictureUpdateDTO

	validateModelFields := validator.New()
	c.BindJSON(&profilePictureUpdate)
	modelFieldsValidationError := validateModelFields.Struct(profilePictureUpdate)

	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	imageUrl := <-storage.UploadFile(profilePictureUpdate.Image, profilePictureUpdate.Name, c)
	newProfilePicture := userModels.UserProfilePicture{
		Id:          profilePictureUpdate.Id,
		UserId:      profilePictureUpdate.UserId,
		Uri:         imageUrl,
		Name:        profilePictureUpdate.Name,
		FileType:    profilePictureUpdate.FileType,
		ContentType: profilePictureUpdate.ContentType,
		Size:        profilePictureUpdate.Size,
	}
	isProfilePictureUpdated := userRepo.UpdateUserProfilePicture(&newProfilePicture)
	if !isProfilePictureUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "profile picture update failed"})
		return
	}
	updatedUser := userRepo.GetUserById(userId)
	if updatedUser == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(updatedUser, updatedUser.AccessToken)})
}

func DeleteUserProfilePicture(c *gin.Context) {
	profilePictureId := c.Param("id")

	client := c.MustGet("user").(*userModels.User)
	user := userRepo.GetUserByEmail(client.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}

	<-storage.DeleteFile(user.ProfilePicture.Name, c)
	isProfilePictureDeleted := userRepo.DeleteUserProfilePicture(profilePictureId)
	if !isProfilePictureDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "profile picture removal failed"})
		return
	}
	updatedUser := userRepo.GetUserById(generalUtilities.ConvertIntToString(user.Id))
	if updatedUser == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(updatedUser, updatedUser.AccessToken)})
}
