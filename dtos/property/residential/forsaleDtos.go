package residential

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type ResidentialPropertyForSaleCreationDto struct {
	ManagerId              int                                                                `json:"managerId"`
	NumberOfRooms          int                                                                `json:"numberOfRooms"`
	Price                  int                                                                `json:"price"`
	NumberOfGarages        int                                                                `json:"numberOfGarages"`
	SizeNumber             int                                                                `json:"sizeNumber"`
	Bedrooms               int                                                                `json:"bedrooms"`
	Bathrooms              int                                                                `json:"bathrooms"`
	YearBuilt              int                                                                `json:"yearBuilt"`
	Storeys                int                                                                `json:"storeys"`
	HasSwimmingPool        bool                                                               `json:"hasSwimmingPool"`
	HasElectricity         bool                                                               `json:"hasElectricity"`
	HasWater               bool                                                               `json:"hasWater"`
	IsNegotiable           bool                                                               `json:"isNegotiable"`
	IsPaved                bool                                                               `json:"isPaved"`
	HasBoreHole            bool                                                               `json:"hasBoreHole"`
	IsPlustered            bool                                                               `json:"isPlustered"`
	IsPainted              bool                                                               `json:"isPainted"`
	IsTiled                bool                                                               `json:"isTiled"`
	HasCeiling             bool                                                               `json:"hasCeiling"`
	Currency               string                                                             `json:"currency"`
	TypeOfExteriorSecurity string                                                             `json:"typeOfExteriorSecurity"`
	Status                 string                                                             `json:"status"`
	MarketingStatement     string                                                             `json:"marketingStatement"`
	Type                   string                                                             `json:"type"`
	SizeDimensions         string                                                             `json:"sizeDimensions"`
	OtherInteriorFeatures  []string                                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                                           `json:"otherExteriorFeatures"`
	PropertyLocation       locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type ResidentialPropertyForSaleResponseDto struct {
	Id                     int                                                `json:"id"`
	UniqueId               int                                                `json:"uniqueId"`
	ManagerId              int                                                `json:"managerId"`
	NumberOfRooms          int                                                `json:"numberOfRooms"`
	Price                  int                                                `json:"price"`
	NumberOfGarages        int                                                `json:"numberOfGarages"`
	SizeNumber             int                                                `json:"sizeNumber"`
	Bedrooms               int                                                `json:"bedrooms"`
	Bathrooms              int                                                `json:"bathrooms"`
	YearBuilt              int                                                `json:"yearBuilt"`
	Storeys                int                                                `json:"storeys"`
	HasSwimmingPool        bool                                               `json:"hasSwimmingPool"`
	HasElectricity         bool                                               `json:"hasElectricity"`
	HasWater               bool                                               `json:"hasWater"`
	IsNegotiable           bool                                               `json:"isNegotiable"`
	IsFavorite             bool                                               `json:"isFavorite"`
	IsPaved                bool                                               `json:"isPaved"`
	HasBoreHole            bool                                               `json:"hasBoreHole"`
	IsPlustered            bool                                               `json:"isPlustered"`
	IsPainted              bool                                               `json:"isPainted"`
	IsTiled                bool                                               `json:"isTiled"`
	HasCeiling             bool                                               `json:"hasCeiling"`
	Currency               string                                             `json:"currency"`
	TypeOfExteriorSecurity string                                             `json:"typeOfExteriorSecurity"`
	Status                 string                                             `json:"status"`
	MarketingStatement     string                                             `json:"marketingStatement"`
	Type                   string                                             `json:"type"`
	SizeDimensions         string                                             `json:"sizeDimensions"`
	OtherInteriorFeatures  []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                           `json:"otherExteriorFeatures"`
	PostedTime             string                                             `json:"postedTime"`
	PropertyLocation       locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights               insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type ResidentialPropertyForSaleWithManagerResponseDto struct {
	Id                     int                                                `json:"id"`
	UniqueId               int                                                `json:"uniqueId"`
	ManagerId              int                                                `json:"managerId"`
	NumberOfRooms          int                                                `json:"numberOfRooms"`
	Price                  int                                                `json:"price"`
	NumberOfGarages        int                                                `json:"numberOfGarages"`
	SizeNumber             int                                                `json:"sizeNumber"`
	Bedrooms               int                                                `json:"bedrooms"`
	Bathrooms              int                                                `json:"bathrooms"`
	YearBuilt              int                                                `json:"yearBuilt"`
	Storeys                int                                                `json:"storeys"`
	HasSwimmingPool        bool                                               `json:"hasSwimmingPool"`
	HasElectricity         bool                                               `json:"hasElectricity"`
	IsFavorite             bool                                               `json:"isFavorite"`
	HasWater               bool                                               `json:"hasWater"`
	IsNegotiable           bool                                               `json:"isNegotiable"`
	IsPaved                bool                                               `json:"isPaved"`
	HasBoreHole            bool                                               `json:"hasBoreHole"`
	IsPlustered            bool                                               `json:"isPlustered"`
	IsPainted              bool                                               `json:"isPainted"`
	IsTiled                bool                                               `json:"isTiled"`
	HasCeiling             bool                                               `json:"hasCeiling"`
	Currency               string                                             `json:"currency"`
	TypeOfExteriorSecurity string                                             `json:"typeOfExteriorSecurity"`
	Status                 string                                             `json:"status"`
	MarketingStatement     string                                             `json:"marketingStatement"`
	Type                   string                                             `json:"type"`
	SizeDimensions         string                                             `json:"sizeDimensions"`
	OtherInteriorFeatures  []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                           `json:"otherExteriorFeatures"`
	PostedTime             string                                             `json:"postedTime"`
	PropertyLocation       locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights               insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager                managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type ResidentialPropertyForSaleUpdateDto struct {
	Id                     int      `json:"id"`
	UniqueId               int      `json:"uniqueId"`
	ManagerId              int      `json:"managerId"`
	NumberOfRooms          int      `json:"numberOfRooms"`
	Price                  int      `json:"price"`
	NumberOfGarages        int      `json:"numberOfGarages"`
	SizeNumber             int      `json:"sizeNumber"`
	Bedrooms               int      `json:"bedrooms"`
	Bathrooms              int      `json:"bathrooms"`
	YearBuilt              int      `json:"yearBuilt"`
	Storeys                int      `json:"storeys"`
	HasSwimmingPool        bool     `json:"hasSwimmingPool"`
	HasElectricity         bool     `json:"hasElectricity"`
	HasWater               bool     `json:"hasWater"`
	IsNegotiable           bool     `json:"isNegotiable"`
	IsPaved                bool     `json:"isPaved"`
	HasBoreHole            bool     `json:"hasBoreHole"`
	IsPlustered            bool     `json:"isPlustered"`
	IsPainted              bool     `json:"isPainted"`
	IsTiled                bool     `json:"isTiled"`
	HasCeiling             bool     `json:"hasCeiling"`
	Currency               string   `json:"currency"`
	TypeOfExteriorSecurity string   `json:"typeOfExteriorSecurity"`
	Status                 string   `json:"status"`
	MarketingStatement     string   `json:"marketingStatement"`
	Type                   string   `json:"type"`
	SizeDimensions         string   `json:"sizeDimensions"`
	OtherInteriorFeatures  []string `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string `json:"otherExteriorFeatures"`
}
