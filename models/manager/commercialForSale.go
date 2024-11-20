package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type CommercialForSaleProperty struct {
	baseModel.MyModel
	Id                    int                                   `json:"id" gorm:"primary_key"`
	ManagerId             int                                   `json:"managerId"`
	UniqueId              int                                   `json:"uniqueId"`
	NumberOfRooms         int                                   `json:"numberOfRooms"`
	Price                 int                                   `json:"price"`
	SizeNumber            int                                   `json:"sizeNumber"`
	YearBuilt             int                                   `json:"yearBuilt"`
	Storeys               int                                   `json:"storeys"`
	HasElectricity        bool                                  `json:"hasElectricity"`
	IsNegotiable          bool                                  `json:"isNegotiable"`
	HasWater              bool                                  `json:"hasWater"`
	Status                string                                `json:"status"`
	Type                  string                                `json:"type"`
	Currency              string                                `json:"currency"`
	SizeDimensions        string                                `json:"sizeDimensions"`
	MarketingStatement    string                                `gorm:"type:text"`
	OtherInteriorFeatures []string                              `gorm:"serializer:json"`
	OtherExteriorFeatures []string                              `gorm:"serializer:json"`
	Manager               Manager                               `json:"manager"`
	Location              propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights      propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia         []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyReport        []propertyModels.PropertyReport       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
