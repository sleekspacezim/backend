package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type CommercialRentalProperty struct {
	baseModel.MyModel
	Id                    int                                   `json:"id" gorm:"primary_key"`
	ManagerId             int                                   `json:"managerId"`
	UniqueId              int                                   `json:"uniqueId"`
	NumberOfRooms         int                                   `json:"numberOfRooms"`
	NumberOfRoomsToLet    int                                   `json:"numberOfRoomsToLet"`
	RentAmount            int                                   `json:"rentAmount"`
	SizeNumber            int                                   `json:"sizeNumber"`
	Storeys               int                                   `json:"storeys"`
	YearBuilt             int                                   `json:"yearBuilt"`
	IsFullSpace           bool                                  `json:"isFullSpace"`
	HasElectricity        bool                                  `json:"hasElectricity"`
	HasWater              bool                                  `json:"hasWater"`
	SizeDimensions        string                                `json:"sizeDimensions"`
	Status                string                                `json:"status"`
	Type                  string                                `json:"type"`
	Currency              string                                `json:"currency"`
	MarketingStatement    string                                `gorm:"type:text"`
	TenantRequirements    []string                              `gorm:"serializer:json"`
	OtherInteriorFeatures []string                              `gorm:"serializer:json"`
	OtherExteriorFeatures []string                              `gorm:"serializer:json"`
	Manager               Manager                               `json:"manager"`
	Location              propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights      propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia         []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyReport        []propertyModels.PropertyReport       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
