package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type Stand struct {
	baseModel.MyModel
	Id                 int                                   `json:"id" gorm:"primary_key"`
	ManagerId          int                                   `json:"managerId"`
	UniqueId           int                                   `json:"uniqueId"`
	Price              int                                   `json:"price"`
	SizeNumber         int                                   `json:"sizeNumber"`
	AreaHasElectricity bool                                  `json:"areaHasElectricity"`
	IsServiced         bool                                  `json:"isServiced"`
	IsNegotiable       bool                                  `json:"isNegotiable"`
	Status             string                                `json:"status"`
	Level              string                                `json:"level"`
	SizeDimensions     string                                `json:"sizeDimensions"`
	Type               string                                `json:"type"`
	Currency           string                                `json:"currency"`
	MarketingStatement string                                `gorm:"type:text"`
	OtherDetails       []string                              `gorm:"serializer:json"`
	Manager            Manager                               `json:"manager"`
	Location           propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights   propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia      []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyReport     []propertyModels.PropertyReport       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
