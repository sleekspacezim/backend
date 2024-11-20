package land

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type LandForSalePropertyCreationDto struct {
	ManagerId          int                                                                `json:"managerId"`
	Price              int                                                                `json:"price"`
	SizeNumber         int                                                                `json:"sizeNumber"`
	AreaHasElectricity bool                                                               `json:"areaHasElectricity"`
	HasWater           bool                                                               `json:"hasWater"`
	IsNegotiable       bool                                                               `json:"isNegotiable"`
	Status             string                                                             `json:"status"`
	Currency           string                                                             `json:"currency"`
	Type               string                                                             `json:"type"`
	MarketingStatement string                                                             `json:"marketingStatement"`
	OtherDetails       []string                                                           `json:"otherDetails"`
	SizeDimensions     string                                                             `json:"sizeDimensions"`
	PropertyLocation   locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media              []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type LandForSalePropertyResponseDto struct {
	Id                 int                                                `json:"id"`
	ManagerId          int                                                `json:"managerId"`
	UniqueId           int                                                `json:"uniqueId"`
	Price              int                                                `json:"price"`
	SizeNumber         int                                                `json:"sizeNumber"`
	AreaHasElectricity bool                                               `json:"areaHasElectricity"`
	IsNegotiable       bool                                               `json:"isNegotiable"`
	IsFavorite         bool                                               `json:"isFavorite"`
	HasWater           bool                                               `json:"hasWater"`
	Status             string                                             `json:"status"`
	Currency           string                                             `json:"currency"`
	MarketingStatement string                                             `json:"marketingStatement"`
	Type               string                                             `json:"type"`
	SizeDimensions     string                                             `json:"sizeDimensions"`
	OtherDetails       []string                                           `json:"otherDetails"`
	PostedTime         string                                             `json:"postedTime"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type LandForSalePropertyWithManagerResponseDto struct {
	Id                 int                                                `json:"id"`
	ManagerId          int                                                `json:"managerId"`
	UniqueId           int                                                `json:"uniqueId"`
	Price              int                                                `json:"price"`
	SizeNumber         int                                                `json:"sizeNumber"`
	AreaHasElectricity bool                                               `json:"areaHasElectricity"`
	IsNegotiable       bool                                               `json:"isNegotiable"`
	IsFavorite         bool                                               `json:"isFavorite"`
	HasWater           bool                                               `json:"hasWater"`
	Status             string                                             `json:"status"`
	Currency           string                                             `json:"currency"`
	MarketingStatement string                                             `json:"marketingStatement"`
	Type               string                                             `json:"type"`
	OtherDetails       []string                                           `json:"otherDetails"`
	SizeDimensions     string                                             `json:"sizeDimensions"`
	PostedTime         string                                             `json:"postedTime"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager            managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type LandForSalePropertyUpdateDto struct {
	Id                 int      `json:"id"`
	ManagerId          int      `json:"managerId"`
	UniqueId           int      `json:"uniqueId"`
	Price              int      `json:"price"`
	SizeNumber         int      `json:"sizeNumber"`
	AreaHasElectricity bool     `json:"areaHasElectricity"`
	IsNegotiable       bool     `json:"isNegotiable"`
	HasWater           bool     `json:"hasWater"`
	Currency           string   `json:"currency"`
	MarketingStatement string   `json:"marketingStatement"`
	Status             string   `json:"status"`
	Type               string   `json:"type"`
	SizeDimensions     string   `json:"sizeDimensions"`
	OtherDetails       []string `json:"otherDetails"`
}
