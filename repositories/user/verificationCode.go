package repositories

import (
	"SleekSpace/db"
	userModels "SleekSpace/models/user"
)

func CreateVerificationCode(user *userModels.User, verificationCode userModels.VerificationCode) bool {
	db.DB.Model(user).Association("RegistrationCode").Replace(verificationCode)
	return true
}

func GetVerificationCodeByUserId(userId string) userModels.VerificationCode {
	var user userModels.User
	err := db.DB.Preload("RegistrationCode").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.RegistrationCode
}

func GetVerificationCodeById(id string) userModels.VerificationCode {
	var code userModels.VerificationCode
	err := db.DB.First(&code, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return code
}

func UpdateVerificationCode(verificationCodeUpdate *userModels.VerificationCode) bool {
	db.DB.Save(verificationCodeUpdate)
	return true
}

func AllVerificationCodes() []userModels.VerificationCode {
	var codes []userModels.VerificationCode
	err := db.DB.Find(&codes)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return codes
}

func DeleteVerficationCode(userId int) bool {
	db.DB.Unscoped().Delete(&userModels.VerificationCode{}, userId)
	return true
}
