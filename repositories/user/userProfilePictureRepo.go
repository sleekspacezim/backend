package repositories

import (
	"SleekSpace/db"
	userModels "SleekSpace/models/user"
)

func CreateUserProfilePicture(user *userModels.User, profilePicture *userModels.UserProfilePicture) bool {
	err := db.DB.Model(user).Association("ProfilePicture").Append(profilePicture)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetAllUsersProfilePictures() []userModels.UserProfilePicture {
	var pictures = []userModels.UserProfilePicture{}
	err := db.DB.Find(&pictures)
	if err != nil {
		println(err.Error, err.Name())
	}
	return pictures
}

func GetUserProfilePictureByManagerId(userId string) userModels.UserProfilePicture {
	var user userModels.User
	err := db.DB.Preload("ProfilePicture").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.ProfilePicture
}

func GetUserProfilePictureById(pictureId string) userModels.UserProfilePicture {
	var profilePicture = userModels.UserProfilePicture{}
	db.DB.First(&profilePicture, pictureId)
	return profilePicture
}

func UpdateUserProfilePicture(profilePictureUpdate *userModels.UserProfilePicture) bool {
	db.DB.Save(profilePictureUpdate)
	return true
}

func DeleteUserProfilePicture(profilePictureId string) bool {
	db.DB.Unscoped().Delete(&userModels.UserProfilePicture{}, profilePictureId)
	return true
}
