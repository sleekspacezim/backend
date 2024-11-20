package repositories

import (
	"SleekSpace/db"
	userModels "SleekSpace/models/user"

	"gorm.io/gorm"
)

func CreateContactNumber(user *userModels.User, contactNumber *userModels.ContactNumber) bool {
	err := db.DB.Model(user).Association("ContactNumbers").Append(contactNumber)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetAllUsersContactNumbers() []userModels.ContactNumber {
	var numbers = []userModels.ContactNumber{}
	err := db.DB.Find(&numbers)
	if err != nil {
		println(err.Error, err.Name())
	}
	return numbers
}

func GetUserContactNumbersByUserId(userId int) []userModels.ContactNumber {
	var user = userModels.User{}
	err := db.DB.Preload("ContactNumbers").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.ContactNumbers
}

func UpdateUserContactNumbers(user *userModels.User, updateContactNumbersList []userModels.ContactNumber) bool {
	user.ContactNumbers = updateContactNumbersList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}
