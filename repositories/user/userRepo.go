package repositories

import (
	"errors"

	"SleekSpace/db"
	"SleekSpace/models/manager"
	userModels "SleekSpace/models/user"
	managerRepo "SleekSpace/repositories/manager"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateUser(user *userModels.User) bool {
	var existingUser userModels.User
	email := user.Email
	result := db.DB.Where("email =?", email).First(&existingUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.DB.Create(user)
		return true
	} else {
		return false
	}
}

func GetUsers() []userModels.User {
	var users = []userModels.User{}
	err := db.DB.Find(&users)
	if err != nil {
		println(err.Error, err.Name())
	}
	return users
}

func GetUserById(id string) *userModels.User {
	var user userModels.User
	result := db.DB.Preload("ContactNumbers").Preload("Location").Preload("ProfilePicture").First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func GetUserAndAllAssociationsById(id string) *userModels.User {
	var user userModels.User
	result := db.DB.Preload(clause.Associations).First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func GetUserByIdWithManager(id string) *userModels.User {
	var user userModels.User
	result := db.DB.Preload("Manager").First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func SaveUserUpdate(update *userModels.User) bool {
	db.DB.Save(update)
	return true
}

func DeleteUserById(id string) bool {
	var user userModels.User
	db.DB.First(&user, id)
	db.DB.Select(clause.Associations).Unscoped().Delete(&user)

	return true
}

func DeleteUserAndCascadeById(user userModels.User, manager manager.Manager) bool {
	isManagerDeleted := managerRepo.DeleteManagerById(&manager)
	if !isManagerDeleted || isManagerDeleted {
		db.DB.Select(clause.Associations).Unscoped().Delete(&user)
		return true
	}
	return true
}

func GetUserByEmail(email string) *userModels.User {
	var user userModels.User
	result := db.DB.Where("email =?", email).Preload("ContactNumbers").Preload("Location").Preload("ProfilePicture").First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return &user
	}
}

func GetUserWithRegistrationCodeById(id string) userModels.User {
	var user userModels.User
	err := db.DB.Preload("RegistrationCode").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user
}
