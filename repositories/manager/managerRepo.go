package repositories

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateManager(manager *managerModels.Manager) bool {
	db.DB.Create(&manager)
	return true

}

func GetManagerByManagerId(managerId string) *managerModels.Manager {
	var manager managerModels.Manager
	result := db.DB.Preload(clause.Associations).First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &manager
}

func GetManagerWithProfilePictureAndContactsByManagerId(managerId string) *managerModels.Manager {
	var manager managerModels.Manager
	result := db.DB.Preload("ProfilePicture").
		Preload("ManagerContactNumbers").
		First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &manager
}

func UpdateManager(managerUpdate *managerModels.Manager) bool {
	db.DB.Save(managerUpdate)
	return true
}

func GetAllManagersContacts() []managerModels.ManagerContactNumber {
	var contacts = []managerModels.ManagerContactNumber{}
	err := db.DB.Find(&contacts)
	if err != nil {
		println(err.Error, err.Name())
	}
	return contacts
}

func GetAllManagers() []managerModels.Manager {
	var managers = []managerModels.Manager{}
	err := db.DB.Find(&managers)
	if err != nil {
		println(err.Error, err.Name())
	}
	return managers
}

func DeleteManagerById(manager *managerModels.Manager) bool {
	db.DB.Select(clause.Associations).Unscoped().Delete(&manager)
	return true
}
