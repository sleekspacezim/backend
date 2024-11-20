package models

import baseModel "SleekSpace/models"

type ContactNumber struct {
	baseModel.MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"userId" gorm:"column:user_id"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
