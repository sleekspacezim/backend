package models

import baseModel "SleekSpace/models"

type PropertyReport struct {
	baseModel.MyModel
	Id             int    `json:"id" gorm:"primary_key"`
	PropertyId     int    `json:"propertyId"`
	ManagerId      int    `json:"managerId"`
	ReporterUserId int    `json:"reporterUserId"`
	Report         string `json:"report" gorm:"type:text"`
}
