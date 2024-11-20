package repositories

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
)

func CreateManagerProfilePicture(manager *managerModels.Manager, profilePicture *managerModels.ManagerProfilePicture) bool {
	err := db.DB.Model(manager).Association("ProfilePicture").Append(profilePicture)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetAllManagersProfilePictures() []managerModels.ManagerProfilePicture {
	var pictures = []managerModels.ManagerProfilePicture{}
	err := db.DB.Find(&pictures)
	if err != nil {
		println(err.Error, err.Name())
	}
	return pictures
}

func GetManagerProfilePictureByManagerId(managerId string) managerModels.ManagerProfilePicture {
	var manager managerModels.Manager
	err := db.DB.Preload("ProfilePicture").First(&manager, managerId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return manager.ProfilePicture
}

func GetProfilePictureById(pictureId string) managerModels.ManagerProfilePicture {
	var profilePicture = managerModels.ManagerProfilePicture{}
	db.DB.First(&profilePicture, pictureId)
	return profilePicture
}

func UpdateProfilePicture(profilePictureUpdate *managerModels.ManagerProfilePicture) bool {
	db.DB.Save(profilePictureUpdate)
	return true
}

func DeleteProfilePicture(profilePictureId string) bool {
	db.DB.Unscoped().Delete(&managerModels.ManagerProfilePicture{}, profilePictureId)
	return true
}
