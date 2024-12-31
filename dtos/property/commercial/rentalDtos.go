package commercial

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type CommercialForRentPropertyCreationDto struct {
	ManagerId             int                                                                `json:"managerId"`
	UniqueId              int                                                                `json:"uniqueId"`
	NumberOfRooms         int                                                                `json:"numberOfRooms"`
	NumberOfRoomsToLet    int                                                                `json:"numberOfRoomsToLet"`
	RentAmount            int                                                                `json:"rentAmount"`
	SizeNumber            int                                                                `json:"sizeNumber"`
	YearBuilt             int                                                                `json:"yearBuilt"`
	Storeys               int                                                                `json:"storeys"`
	IsFullSpace           bool                                                               `json:"isFullSpace"`
	HasElectricity        bool                                                               `json:"hasElectricity"`
	HasWater              bool                                                               `json:"hasWater"`
	Currency              string                                                             `json:"currency"`
	SizeDimensions        string                                                             `json:"sizeDimensions"`
	Status                string                                                             `json:"status"`
	Type                  string                                                             `json:"type"`
	MarketingStatement    string                                                             `json:"marketingStatement"`
	TenantRequirements    []string                                                           `json:"tenantRequirements"`
	OtherInteriorFeatures []string                                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type CommercialForRentPropertyResponseDto struct {
	Id                    int                                                `json:"id"`
	PostedTime            string                                             `json:"postedTime"`
	ManagerId             int                                                `json:"managerId"`
	UniqueId              int                                                `json:"uniqueId"`
	NumberOfRooms         int                                                `json:"numberOfRooms"`
	NumberOfRoomsToLet    int                                                `json:"numberOfRoomsToLet"`
	RentAmount            int                                                `json:"rentAmount"`
	SizeNumber            int                                                `json:"sizeNumber"`
	YearBuilt             int                                                `json:"yearBuilt"`
	Storeys               int                                                `json:"storeys"`
	IsFullSpace           bool                                               `json:"isFullSpace"`
	HasElectricity        bool                                               `json:"hasElectricity"`
	IsFavorite            bool                                               `json:"isFavorite"`
	HasWater              bool                                               `json:"hasWater"`
	Currency              string                                             `json:"currency"`
	SizeDimensions        string                                             `json:"sizeDimensions"`
	Status                string                                             `json:"status"`
	Type                  string                                             `json:"type"`
	MarketingStatement    string                                             `json:"marketingStatement"`
	TenantRequirements    []string                                           `json:"tenantRequirements"`
	OtherInteriorFeatures []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights              insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type CommercialForRentPropertyWithManagerResponseDto struct {
	Id                    int                                                `json:"id"`
	ManagerId             int                                                `json:"managerId"`
	UniqueId              int                                                `json:"uniqueId"`
	NumberOfRooms         int                                                `json:"numberOfRooms"`
	NumberOfRoomsToLet    int                                                `json:"numberOfRoomsToLet"`
	RentAmount            int                                                `json:"rentAmount"`
	SizeNumber            int                                                `json:"sizeNumber"`
	YearBuilt             int                                                `json:"yearBuilt"`
	Storeys               int                                                `json:"storeys"`
	IsFullSpace           bool                                               `json:"isFullSpace"`
	HasElectricity        bool                                               `json:"hasElectricity"`
	HasWater              bool                                               `json:"hasWater"`
	IsFavorite            bool                                               `json:"isFavorite"`
	Currency              string                                             `json:"currency"`
	SizeDimensions        string                                             `json:"sizeDimensions"`
	Status                string                                             `json:"status"`
	Type                  string                                             `json:"type"`
	MarketingStatement    string                                             `json:"marketingStatement"`
	PostedTime            string                                             `json:"postedTime"`
	TenantRequirements    []string                                           `json:"tenantRequirements"`
	OtherInteriorFeatures []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights              insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager               managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type CommercialForRentPropertyUpdateDto struct {
	Id                    int      `json:"id"`
	ManagerId             int      `json:"managerId"`
	UniqueId              int      `json:"uniqueId"`
	NumberOfRooms         int      `json:"numberOfRooms"`
	NumberOfRoomsToLet    int      `json:"numberOfRoomsToLet"`
	RentAmount            int      `json:"rentAmount"`
	SizeNumber            int      `json:"sizeNumber"`
	YearBuilt             int      `json:"yearBuilt"`
	Storeys               int      `json:"storeys"`
	IsFullSpace           bool     `json:"isFullSpace"`
	HasElectricity        bool     `json:"hasElectricity"`
	HasWater              bool     `json:"hasWater"`
	Currency              string   `json:"currency"`
	SizeDimensions        string   `json:"sizeDimensions"`
	Status                string   `json:"status"`
	Type                  string   `json:"type"`
	MarketingStatement    string   `json:"marketingStatement"`
	TenantRequirements    []string `json:"tenantRequirements"`
	OtherInteriorFeatures []string `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string `json:"otherExteriorFeatures"`
}
