package insights

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
)

func GetAllPropertyLocations() []propertyModels.PropertyLocation {
	var locations = []propertyModels.PropertyLocation{}
	err := db.DB.Find(&locations)
	if err != nil {
		println(err.Error, err.Name())
	}
	return locations
}

func GetPropertyLocationById(propertyLocationId string) *propertyModels.PropertyLocation {
	var location = propertyModels.PropertyLocation{}
	result := db.DB.First(&location, propertyLocationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &location
}

func UpdatePropertyLocation(propertyLocationUpdate *propertyModels.PropertyLocation) bool {
	db.DB.Save(propertyLocationUpdate)
	return true
}

func DeletePropertyLocation(propertyLocationId string) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyLocation{}, propertyLocationId)
	return true
}
