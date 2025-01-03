package residential

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type ResidentialPropertyForRentCreationDto struct {
	ManagerId              int                                                                `json:"managerId"`
	NumberOfRoomsToLet     int                                                                `json:"numberOfRoomsToLet"`
	NumberOfRooms          int                                                                `json:"numberOfRooms"`
	RentAmount             int                                                                `json:"rentAmount"`
	NumberOfGarages        int                                                                `json:"numberOfGarages"`
	SizeNumber             int                                                                `json:"sizeNumber"`
	Bathrooms              int                                                                `json:"bathrooms"`
	Bedrooms               int                                                                `json:"bedrooms"`
	YearBuilt              int                                                                `json:"yearBuilt"`
	Storeys                int                                                                `json:"storeys"`
	IsPaved                bool                                                               `json:"isPaved"`
	HasBoreHole            bool                                                               `json:"hasBoreHole"`
	IsPlustered            bool                                                               `json:"isPlustered"`
	IsPainted              bool                                                               `json:"isPainted"`
	IsTiled                bool                                                               `json:"isTiled"`
	HasCeiling             bool                                                               `json:"hasCeiling"`
	IsFullHouse            bool                                                               `json:"isFullHouse"`
	HasElectricity         bool                                                               `json:"hasElectricity"`
	HasWater               bool                                                               `json:"hasWater"`
	HasSwimmingPool        bool                                                               `json:"hasSwimmingPool"`
	Currency               string                                                             `json:"currency"`
	TypeOfExteriorSecurity string                                                             `json:"typeOfExteriorSecurity"`
	MarketingStatement     string                                                             `json:"marketingStatement"`
	Status                 string                                                             `json:"status"`
	Type                   string                                                             `json:"type"`
	SizeDimensions         string                                                             `json:"sizeDimensions"`
	TenantRequirements     []string                                                           `json:"tenantRequirements"`
	OtherInteriorFeatures  []string                                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                                           `json:"otherExteriorFeatures"`
	PropertyLocation       locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type ResidentialPropertyForRentResponseDto struct {
	Id                     int                                                `json:"id"`
	ManagerId              int                                                `json:"managerId"`
	UniqueId               int                                                `json:"uniqueId"`
	NumberOfRoomsToLet     int                                                `json:"numberOfRoomsToLet"`
	NumberOfRooms          int                                                `json:"numberOfRooms"`
	RentAmount             int                                                `json:"rentAmount"`
	NumberOfGarages        int                                                `json:"numberOfGarages"`
	SizeNumber             int                                                `json:"sizeNumber"`
	Bathrooms              int                                                `json:"bathrooms"`
	Storeys                int                                                `json:"storeys"`
	Bedrooms               int                                                `json:"bedrooms"`
	YearBuilt              int                                                `json:"yearBuilt"`
	IsPaved                bool                                               `json:"isPaved"`
	HasBoreHole            bool                                               `json:"hasBoreHole"`
	IsPlustered            bool                                               `json:"isPlustered"`
	IsPainted              bool                                               `json:"isPainted"`
	IsTiled                bool                                               `json:"isTiled"`
	HasCeiling             bool                                               `json:"hasCeiling"`
	IsFullHouse            bool                                               `json:"isFullHouse"`
	HasElectricity         bool                                               `json:"hasElectricity"`
	HasWater               bool                                               `json:"hasWater"`
	HasSwimmingPool        bool                                               `json:"hasSwimmingPool"`
	IsFavorite             bool                                               `json:"isFavorite"`
	Currency               string                                             `json:"currency"`
	TypeOfExteriorSecurity string                                             `json:"typeOfExteriorSecurity"`
	MarketingStatement     string                                             `json:"marketingStatement"`
	Status                 string                                             `json:"status"`
	Type                   string                                             `json:"type"`
	SizeDimensions         string                                             `json:"sizeDimensions"`
	TenantRequirements     []string                                           `json:"tenantRequirements"`
	OtherInteriorFeatures  []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                           `json:"otherExteriorFeatures"`
	PostedTime             string                                             `json:"postedTime"`
	PropertyLocation       locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights               insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type ResidentialPropertyForRentWithManagerResponseDto struct {
	Id                     int                                                `json:"id"`
	ManagerId              int                                                `json:"managerId"`
	UniqueId               int                                                `json:"uniqueId"`
	NumberOfRoomsToLet     int                                                `json:"numberOfRoomsToLet"`
	NumberOfRooms          int                                                `json:"numberOfRooms"`
	RentAmount             int                                                `json:"rentAmount"`
	NumberOfGarages        int                                                `json:"numberOfGarages"`
	SizeNumber             int                                                `json:"sizeNumber"`
	Bathrooms              int                                                `json:"bathrooms"`
	Bedrooms               int                                                `json:"bedrooms"`
	Storeys                int                                                `json:"storeys"`
	YearBuilt              int                                                `json:"yearBuilt"`
	IsPaved                bool                                               `json:"isPaved"`
	HasBoreHole            bool                                               `json:"hasBoreHole"`
	IsPlustered            bool                                               `json:"isPlustered"`
	IsPainted              bool                                               `json:"isPainted"`
	IsTiled                bool                                               `json:"isTiled"`
	HasCeiling             bool                                               `json:"hasCeiling"`
	IsFullHouse            bool                                               `json:"isFullHouse"`
	HasElectricity         bool                                               `json:"hasElectricity"`
	HasWater               bool                                               `json:"hasWater"`
	HasSwimmingPool        bool                                               `json:"hasSwimmingPool"`
	IsFavorite             bool                                               `json:"isFavorite"`
	Currency               string                                             `json:"currency"`
	TypeOfExteriorSecurity string                                             `json:"typeOfExteriorSecurity"`
	MarketingStatement     string                                             `json:"marketingStatement"`
	Status                 string                                             `json:"status"`
	Type                   string                                             `json:"type"`
	SizeDimensions         string                                             `json:"sizeDimensions"`
	TenantRequirements     []string                                           `json:"tenantRequirements"`
	OtherInteriorFeatures  []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string                                           `json:"otherExteriorFeatures"`
	PostedTime             string                                             `json:"postedTime"`
	PropertyLocation       locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights               insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                  []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager                managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type ResidentialPropertyForRentUpdateDto struct {
	Id                     int      `json:"id"`
	ManagerId              int      `json:"managerId"`
	UniqueId               int      `json:"uniqueId"`
	NumberOfRoomsToLet     int      `json:"numberOfRoomsToLet"`
	NumberOfRooms          int      `json:"numberOfRooms"`
	RentAmount             int      `json:"rentAmount"`
	NumberOfGarages        int      `json:"numberOfGarages"`
	SizeNumber             int      `json:"sizeNumber"`
	YearBuilt              int      `json:"yearBuilt"`
	Bathrooms              int      `json:"bathrooms"`
	Storeys                int      `json:"storeys"`
	Bedrooms               int      `json:"bedrooms"`
	IsPaved                bool     `json:"isPaved"`
	HasBoreHole            bool     `json:"hasBoreHole"`
	IsPlustered            bool     `json:"isPlustered"`
	IsPainted              bool     `json:"isPainted"`
	IsTiled                bool     `json:"isTiled"`
	HasCeiling             bool     `json:"hasCeiling"`
	IsFullHouse            bool     `json:"isFullHouse"`
	HasElectricity         bool     `json:"hasElectricity"`
	HasWater               bool     `json:"hasWater"`
	HasSwimmingPool        bool     `json:"hasSwimmingPool"`
	Currency               string   `json:"currency"`
	TypeOfExteriorSecurity string   `json:"typeOfExteriorSecurity"`
	MarketingStatement     string   `json:"marketingStatement"`
	Status                 string   `json:"status"`
	Type                   string   `json:"type"`
	SizeDimensions         string   `json:"sizeDimensions"`
	TenantRequirements     []string `json:"tenantRequirements"`
	OtherInteriorFeatures  []string `json:"otherInteriorFeatures"`
	OtherExteriorFeatures  []string `json:"otherExteriorFeatures"`
}
