package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type ResidentialPropertyForSale struct {
	baseModel.MyModel
	Id                     int                                   `json:"id" gorm:"primary_key"`
	ManagerId              int                                   `json:"managerId"`
	UniqueId               int                                   `json:"uniqueId"`
	NumberOfRooms          int                                   `json:"numberOfRooms"`
	Price                  int                                   `json:"price"`
	NumberOfGarages        int                                   `json:"numberOfGarages"`
	SizeNumber             int                                   `json:"sizeNumber"`
	Storeys                int                                   `json:"storeys"`
	Bedrooms               int                                   `json:"bedrooms"`
	Bathrooms              int                                   `json:"bathrooms"`
	YearBuilt              int                                   `json:"yearBuilt"`
	HasSwimmingPool        bool                                  `json:"hasSwimmingPool"`
	HasElectricity         bool                                  `json:"hasElectricity"`
	HasWater               bool                                  `json:"hasWater"`
	IsPaved                bool                                  `json:"isPaved"`
	HasBoreHole            bool                                  `json:"hasBoreHole"`
	IsPlustered            bool                                  `json:"isPlustered"`
	IsPainted              bool                                  `json:"isPainted"`
	IsTiled                bool                                  `json:"isTiled"`
	HasCeiling             bool                                  `json:"hasCeiling"`
	IsNegotiable           bool                                  `json:"isNegotiable"`
	TypeOfExteriorSecurity string                                `json:"typeOfExteriorSecurity"`
	Status                 string                                `json:"status"`
	Type                   string                                `json:"type"`
	SizeDimensions         string                                `json:"sizeDimensions"`
	Currency               string                                `json:"currency"`
	MarketingStatement     string                                `gorm:"type:text"`
	OtherInteriorFeatures  []string                              `gorm:"serializer:json"`
	OtherExteriorFeatures  []string                              `gorm:"serializer:json"`
	Manager                Manager                               `json:"manager"`
	Location               propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights       propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia          []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyReport         []propertyModels.PropertyReport       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
