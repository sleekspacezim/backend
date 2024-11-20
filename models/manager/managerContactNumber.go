package manager

import baseModel "SleekSpace/models"

type ManagerContactNumber struct {
	baseModel.MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	ManagerId    int    `json:"managerId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
