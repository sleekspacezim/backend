package models

import baseModel "SleekSpace/models"

type PropertyInsights struct {
	baseModel.MyModel
	Id                int    `json:"id" gorm:"primary_key"`
	PropertyId        int    `json:"propertyId"`
	Views             int    `json:"views"`
	AddedToFavourites int    `json:"addedToFavourites"`
	Shared            int    `json:"shared"`
	ContactInfoViews  int    `json:"contactInfoViews"`
	PropertyType      string `json:"propertyType"`
}
