package insights

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
)

func GetAllPropertiesInsights() []propertyModels.PropertyInsights {
	var insightsList = []propertyModels.PropertyInsights{}
	err := db.DB.Find(&insightsList)
	if err != nil {
		println(err.Error, err.Name())
	}
	return insightsList
}

func GetPropertyInsightsByPropertyId(propertyId string) *propertyModels.PropertyInsights {
	var insights = propertyModels.PropertyInsights{}
	result := db.DB.Where("property_id =?", propertyId).First(&insights)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &insights
}

func GetPropertyInsightsById(propertyInsightsId string) *propertyModels.PropertyInsights {
	var insights = propertyModels.PropertyInsights{}
	result := db.DB.First(&insights, propertyInsightsId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &insights
}

func UpdatePropertyInsights(propertyInsightsUpdate *propertyModels.PropertyInsights) bool {
	db.DB.Save(propertyInsightsUpdate)
	return true
}

func DeletePropertyInsights(propertyInsightsId string) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyInsights{}, propertyInsightsId)
	return true
}
