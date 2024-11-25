package property

import (
	residentialPropertyDtos "SleekSpace/dtos/property/residential"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func ResidentialRentalPropertyWithManagerResponse(residentialRentalModel managerModels.ResidentialRentalProperty) residentialPropertyDtos.ResidentialPropertyForRentWithManagerResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForRentWithManagerResponseDto{
		Id:                     residentialRentalModel.Id,
		ManagerId:              residentialRentalModel.ManagerId,
		UniqueId:               residentialRentalModel.UniqueId,
		RentAmount:             residentialRentalModel.RentAmount,
		SizeNumber:             residentialRentalModel.SizeNumber,
		SizeDimensions:         residentialRentalModel.SizeDimensions,
		Currency:               residentialRentalModel.Currency,
		MarketingStatement:     residentialRentalModel.MarketingStatement,
		Status:                 residentialRentalModel.Status,
		Type:                   residentialRentalModel.Type,
		YearBuilt:              residentialRentalModel.YearBuilt,
		Bedrooms:               residentialRentalModel.Bedrooms,
		Bathrooms:              residentialRentalModel.Bathrooms,
		Storeys:                residentialRentalModel.Storeys,
		HasElectricity:         residentialRentalModel.HasElectricity,
		HasWater:               residentialRentalModel.HasWater,
		IsFullHouse:            residentialRentalModel.IsFullHouse,
		NumberOfRoomsToLet:     residentialRentalModel.NumberOfRoomsToLet,
		NumberOfGarages:        residentialRentalModel.NumberOfGarages,
		HasSwimmingPool:        residentialRentalModel.HasSwimmingPool,
		OtherInteriorFeatures:  residentialRentalModel.OtherInteriorFeatures,
		OtherExteriorFeatures:  residentialRentalModel.OtherExteriorFeatures,
		TotalNumberOfRooms:     residentialRentalModel.TotalNumberOfRooms,
		IsPaved:                residentialRentalModel.IsPaved,
		IsPlustered:            residentialRentalModel.IsPlustered,
		IsPainted:              residentialRentalModel.IsPainted,
		IsTiled:                residentialRentalModel.IsTiled,
		IsFavorite:             false,
		HasBoreHole:            residentialRentalModel.HasBoreHole,
		HasCeiling:             residentialRentalModel.HasCeiling,
		TypeOfExteriorSecurity: residentialRentalModel.TypeOfExteriorSecurity,
		TenantRequirements:     residentialRentalModel.TenantRequirements,
		PostedTime:             generalUtilities.GetTimePassed(residentialRentalModel.CreatedAt),
		PropertyLocation:       PropertyLocationResponse(residentialRentalModel.Location),
		Insights:               PropertyInsightsResponse(residentialRentalModel.PropertyInsights),
		Media:                  ProcessedPropertyImageAndVideosListToResponse(residentialRentalModel.PropertyMedia),
		Manager:                managerUtilities.ManagerResponse(&residentialRentalModel.Manager),
	}
}

func ResidentialForSalePropertyWithManagerResponse(residentialForSaleModel managerModels.ResidentialPropertyForSale) residentialPropertyDtos.ResidentialPropertyForSaleWithManagerResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForSaleWithManagerResponseDto{
		Id:                     residentialForSaleModel.Id,
		ManagerId:              residentialForSaleModel.ManagerId,
		UniqueId:               residentialForSaleModel.UniqueId,
		Price:                  residentialForSaleModel.Price,
		SizeNumber:             residentialForSaleModel.SizeNumber,
		SizeDimensions:         residentialForSaleModel.SizeDimensions,
		Status:                 residentialForSaleModel.Status,
		Currency:               residentialForSaleModel.Currency,
		Type:                   residentialForSaleModel.Type,
		YearBuilt:              residentialForSaleModel.YearBuilt,
		MarketingStatement:     residentialForSaleModel.MarketingStatement,
		Bedrooms:               residentialForSaleModel.Bedrooms,
		Bathrooms:              residentialForSaleModel.Bathrooms,
		Storeys:                residentialForSaleModel.Storeys,
		HasElectricity:         residentialForSaleModel.HasElectricity,
		HasWater:               residentialForSaleModel.HasWater,
		NumberOfRooms:          residentialForSaleModel.NumberOfRooms,
		NumberOfGarages:        residentialForSaleModel.NumberOfGarages,
		HasSwimmingPool:        residentialForSaleModel.HasSwimmingPool,
		IsNegotiable:           residentialForSaleModel.IsNegotiable,
		IsPaved:                residentialForSaleModel.IsPaved,
		IsPlustered:            residentialForSaleModel.IsPlustered,
		IsPainted:              residentialForSaleModel.IsPainted,
		IsTiled:                residentialForSaleModel.IsTiled,
		HasBoreHole:            residentialForSaleModel.HasBoreHole,
		HasCeiling:             residentialForSaleModel.HasCeiling,
		IsFavorite:             false,
		TypeOfExteriorSecurity: residentialForSaleModel.TypeOfExteriorSecurity,
		OtherExteriorFeatures:  residentialForSaleModel.OtherExteriorFeatures,
		OtherInteriorFeatures:  residentialForSaleModel.OtherInteriorFeatures,
		PostedTime:             generalUtilities.GetTimePassed(residentialForSaleModel.CreatedAt),
		PropertyLocation:       PropertyLocationResponse(residentialForSaleModel.Location),
		Insights:               PropertyInsightsResponse(residentialForSaleModel.PropertyInsights),
		Media:                  ProcessedPropertyImageAndVideosListToResponse(residentialForSaleModel.PropertyMedia),
		Manager:                managerUtilities.ManagerResponse(&residentialForSaleModel.Manager),
	}
}

func ResidentialForSalePropertyResponse(residentialForSaleModel managerModels.ResidentialPropertyForSale) residentialPropertyDtos.ResidentialPropertyForSaleResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForSaleResponseDto{
		Id:                     residentialForSaleModel.Id,
		ManagerId:              residentialForSaleModel.ManagerId,
		UniqueId:               residentialForSaleModel.UniqueId,
		Price:                  residentialForSaleModel.Price,
		SizeNumber:             residentialForSaleModel.SizeNumber,
		SizeDimensions:         residentialForSaleModel.SizeDimensions,
		Status:                 residentialForSaleModel.Status,
		Type:                   residentialForSaleModel.Type,
		Currency:               residentialForSaleModel.Currency,
		YearBuilt:              residentialForSaleModel.YearBuilt,
		Bedrooms:               residentialForSaleModel.Bedrooms,
		Bathrooms:              residentialForSaleModel.Bathrooms,
		Storeys:                residentialForSaleModel.Storeys,
		HasElectricity:         residentialForSaleModel.HasElectricity,
		HasWater:               residentialForSaleModel.HasWater,
		NumberOfRooms:          residentialForSaleModel.NumberOfRooms,
		MarketingStatement:     residentialForSaleModel.MarketingStatement,
		NumberOfGarages:        residentialForSaleModel.NumberOfGarages,
		HasSwimmingPool:        residentialForSaleModel.HasSwimmingPool,
		IsNegotiable:           residentialForSaleModel.IsNegotiable,
		IsPaved:                residentialForSaleModel.IsPaved,
		IsPlustered:            residentialForSaleModel.IsPlustered,
		IsPainted:              residentialForSaleModel.IsPainted,
		IsTiled:                residentialForSaleModel.IsTiled,
		HasBoreHole:            residentialForSaleModel.HasBoreHole,
		IsFavorite:             false,
		HasCeiling:             residentialForSaleModel.HasCeiling,
		TypeOfExteriorSecurity: residentialForSaleModel.TypeOfExteriorSecurity,
		OtherExteriorFeatures:  residentialForSaleModel.OtherExteriorFeatures,
		OtherInteriorFeatures:  residentialForSaleModel.OtherInteriorFeatures,
		PostedTime:             generalUtilities.GetTimePassed(residentialForSaleModel.CreatedAt),
		PropertyLocation:       PropertyLocationResponse(residentialForSaleModel.Location),
		Insights:               PropertyInsightsResponse(residentialForSaleModel.PropertyInsights),
		Media:                  ProcessedPropertyImageAndVideosListToResponse(residentialForSaleModel.PropertyMedia),
	}
}

func ResidentialRentalPropertyResponse(residentialRentalModel managerModels.ResidentialRentalProperty) residentialPropertyDtos.ResidentialPropertyForRentResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForRentResponseDto{
		Id:                     residentialRentalModel.Id,
		ManagerId:              residentialRentalModel.ManagerId,
		UniqueId:               residentialRentalModel.UniqueId,
		RentAmount:             residentialRentalModel.RentAmount,
		SizeNumber:             residentialRentalModel.SizeNumber,
		SizeDimensions:         residentialRentalModel.SizeDimensions,
		Status:                 residentialRentalModel.Status,
		Currency:               residentialRentalModel.Currency,
		Type:                   residentialRentalModel.Type,
		YearBuilt:              residentialRentalModel.YearBuilt,
		Bedrooms:               residentialRentalModel.Bedrooms,
		Bathrooms:              residentialRentalModel.Bathrooms,
		Storeys:                residentialRentalModel.Storeys,
		HasElectricity:         residentialRentalModel.HasElectricity,
		HasWater:               residentialRentalModel.HasWater,
		IsFullHouse:            residentialRentalModel.IsFullHouse,
		IsFavorite:             false,
		NumberOfRoomsToLet:     residentialRentalModel.NumberOfRoomsToLet,
		MarketingStatement:     residentialRentalModel.MarketingStatement,
		NumberOfGarages:        residentialRentalModel.NumberOfGarages,
		HasSwimmingPool:        residentialRentalModel.HasSwimmingPool,
		OtherExteriorFeatures:  residentialRentalModel.OtherExteriorFeatures,
		OtherInteriorFeatures:  residentialRentalModel.OtherInteriorFeatures,
		TotalNumberOfRooms:     residentialRentalModel.TotalNumberOfRooms,
		IsPaved:                residentialRentalModel.IsPaved,
		IsPlustered:            residentialRentalModel.IsPlustered,
		IsPainted:              residentialRentalModel.IsPainted,
		IsTiled:                residentialRentalModel.IsTiled,
		HasBoreHole:            residentialRentalModel.HasBoreHole,
		HasCeiling:             residentialRentalModel.HasCeiling,
		TypeOfExteriorSecurity: residentialRentalModel.TypeOfExteriorSecurity,
		TenantRequirements:     residentialRentalModel.TenantRequirements,
		PostedTime:             generalUtilities.GetTimePassed(residentialRentalModel.CreatedAt),
		PropertyLocation:       PropertyLocationResponse(residentialRentalModel.Location),
		Insights:               PropertyInsightsResponse(residentialRentalModel.PropertyInsights),
		Media:                  ProcessedPropertyImageAndVideosListToResponse(residentialRentalModel.PropertyMedia),
	}
}