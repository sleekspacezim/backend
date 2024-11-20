package models

import baseModel "SleekSpace/models"

type PropertyLocation struct {
	baseModel.MyModel
	Id           int      `json:"id" gorm:"primary_key"`
	PropertyId   int      `json:"propertyId"`
	DisplayName  string   `json:"displayName"`
	Boundingbox  []string `json:"boundingbox" gorm:"serializer:json"`
	Lat          string   `json:"lat"`
	Lon          string   `json:"lon"`
	Surburb      string   `json:"surburb"`
	City         string   `json:"city"`
	County       string   `json:"county"`
	Province     string   `json:"province"`
	Country      string   `json:"country"`
	CountryCode  string   `json:"countryCode"`
	PropertyType string   `json:"propertyType"`
}
