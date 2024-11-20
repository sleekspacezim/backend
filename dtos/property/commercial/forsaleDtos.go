package commercial

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type CommercialForSalePropertyCreationDto struct {
	ManagerId             int                                                                `json:"managerId"`
	UniqueId              int                                                                `json:"uniqueId"`
	NumberOfRooms         int                                                                `json:"numberOfRooms"`
	Price                 int                                                                `json:"price"`
	SizeNumber            int                                                                `json:"sizeNumber"`
	YearBuilt             int                                                                `json:"yearBuilt"`
	Storeys               int                                                                `json:"storeys"`
	HasElectricity        bool                                                               `json:"hasElectricity"`
	HasWater              bool                                                               `json:"hasWater"`
	IsNegotiable          bool                                                               `json:"isNegotiable"`
	Currency              string                                                             `json:"currency"`
	Status                string                                                             `json:"status"`
	Type                  string                                                             `json:"type"`
	MarketingStatement    string                                                             `json:"marketingStatement"`
	SizeDimensions        string                                                             `json:"sizeDimensions"`
	OtherInteriorFeatures []string                                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type CommercialForSalePropertyResponseDto struct {
	Id                    int                                                `json:"id"`
	ManagerId             int                                                `json:"managerId"`
	UniqueId              int                                                `json:"uniqueId"`
	NumberOfRooms         int                                                `json:"numberOfRooms"`
	Price                 int                                                `json:"price"`
	SizeNumber            int                                                `json:"sizeNumber"`
	YearBuilt             int                                                `json:"yearBuilt"`
	Storeys               int                                                `json:"storeys"`
	HasElectricity        bool                                               `json:"hasElectricity"`
	HasWater              bool                                               `json:"hasWater"`
	IsNegotiable          bool                                               `json:"isNegotiable"`
	IsFavorite            bool                                               `json:"isFavorite"`
	Currency              string                                             `json:"currency"`
	Status                string                                             `json:"status"`
	Type                  string                                             `json:"type"`
	PostedTime            string                                             `json:"postedTime"`
	MarketingStatement    string                                             `json:"marketingStatement"`
	SizeDimensions        string                                             `json:"sizeDimensions"`
	OtherInteriorFeatures []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights              insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type CommercialForSalePropertyWithManagerResponseDto struct {
	Id                    int                                                `json:"id"`
	ManagerId             int                                                `json:"managerId"`
	UniqueId              int                                                `json:"uniqueId"`
	NumberOfRooms         int                                                `json:"numberOfRooms"`
	Price                 int                                                `json:"price"`
	SizeNumber            int                                                `json:"sizeNumber"`
	YearBuilt             int                                                `json:"yearBuilt"`
	Storeys               int                                                `json:"storeys"`
	HasElectricity        bool                                               `json:"hasElectricity"`
	HasWater              bool                                               `json:"hasWater"`
	IsNegotiable          bool                                               `json:"isNegotiable"`
	IsFavorite            bool                                               `json:"isFavorite"`
	PostedTime            string                                             `json:"postedTime"`
	Currency              string                                             `json:"currency"`
	Status                string                                             `json:"status"`
	Type                  string                                             `json:"type"`
	MarketingStatement    string                                             `json:"marketingStatement"`
	SizeDimensions        string                                             `json:"sizeDimensions"`
	OtherInteriorFeatures []string                                           `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string                                           `json:"otherExteriorFeatures"`
	PropertyLocation      locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights              insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media                 []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager               managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type CommercialForSalePropertyUpdateDto struct {
	Id                    int      `json:"id"`
	ManagerId             int      `json:"managerId"`
	UniqueId              int      `json:"uniqueId"`
	NumberOfRooms         int      `json:"numberOfRooms"`
	Price                 int      `json:"price"`
	SizeNumber            int      `json:"sizeNumber"`
	YearBuilt             int      `json:"yearBuilt"`
	Storeys               int      `json:"storeys"`
	HasElectricity        bool     `json:"hasElectricity"`
	HasWater              bool     `json:"hasWater"`
	IsNegotiable          bool     `json:"isNegotiable"`
	Currency              string   `json:"currency"`
	Status                string   `json:"status"`
	Type                  string   `json:"type"`
	MarketingStatement    string   `json:"marketingStatement"`
	SizeDimensions        string   `json:"sizeDimensions"`
	OtherInteriorFeatures []string `json:"otherInteriorFeatures"`
	OtherExteriorFeatures []string `json:"otherExteriorFeatures"`
}
