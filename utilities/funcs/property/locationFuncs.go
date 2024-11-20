package property

import (
	propertyLocationDtos "SleekSpace/dtos/property/location"
	propertyModels "SleekSpace/models/property"
)

func PropertyLocationResponse(location propertyModels.PropertyLocation) propertyLocationDtos.PropertyLocationUpdateAndResponseDto {
	return propertyLocationDtos.PropertyLocationUpdateAndResponseDto{
		Id:           location.Id,
		PropertyId:   location.PropertyId,
		Boundingbox:  location.Boundingbox,
		Lat:          location.Lat,
		Lon:          location.Lon,
		DisplayName:  location.DisplayName,
		City:         location.City,
		County:       location.County,
		Country:      location.Country,
		CountryCode:  location.CountryCode,
		Province:     location.Province,
		Surburb:      location.Surburb,
		PropertyType: location.PropertyType,
	}
}
