package models

import (
	baseModel "SleekSpace/models"
	"time"
)

type VerificationCode struct {
	baseModel.MyModel
	Id         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"userId"`
	Code       int       `json:"code"`
	ExpiryDate time.Time `json:"expirydate"`
}
