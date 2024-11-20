package models

import baseModel "SleekSpace/models"

type Location struct {
	baseModel.MyModel
	Id          int      `json:"id" gorm:"primary_key"`
	UserId      int      `json:"userId"`
	DisplayName string   `json:"displayName"`
	Boundingbox []string `json:"boundingbox" gorm:"serializer:json"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Surburb     string   `json:"surburb"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	Province    string   `json:"province"`
	Country     string   `json:"country"`
	CountryCode string   `json:"countryCode"`
}
