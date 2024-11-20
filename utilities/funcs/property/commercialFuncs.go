package property

import (
	commercialPropertyDtos "SleekSpace/dtos/property/commercial"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func CommercialPropertyForRentResponse(commercialPropertyForRent managerModels.CommercialRentalProperty) commercialPropertyDtos.CommercialForRentPropertyResponseDto {
	return commercialPropertyDtos.CommercialForRentPropertyResponseDto{
		Id:                    commercialPropertyForRent.Id,
		ManagerId:             commercialPropertyForRent.ManagerId,
		UniqueId:              commercialPropertyForRent.UniqueId,
		RentAmount:            commercialPropertyForRent.RentAmount,
		SizeNumber:            commercialPropertyForRent.SizeNumber,
		SizeDimensions:        commercialPropertyForRent.SizeDimensions,
		Status:                commercialPropertyForRent.Status,
		Type:                  commercialPropertyForRent.Type,
		YearBuilt:             commercialPropertyForRent.YearBuilt,
		Storeys:               commercialPropertyForRent.Storeys,
		HasElectricity:        commercialPropertyForRent.HasElectricity,
		HasWater:              commercialPropertyForRent.HasWater,
		MarketingStatement:    commercialPropertyForRent.MarketingStatement,
		IsFavorite:            false,
		NumberOfRooms:         commercialPropertyForRent.NumberOfRooms,
		IsFullSpace:           commercialPropertyForRent.IsFullSpace,
		TenantRequirements:    commercialPropertyForRent.TenantRequirements,
		OtherExteriorFeatures: commercialPropertyForRent.OtherExteriorFeatures,
		OtherInteriorFeatures: commercialPropertyForRent.OtherInteriorFeatures,
		Currency:              commercialPropertyForRent.Currency,
		PostedTime:            generalUtilities.GetTimePassed(commercialPropertyForRent.CreatedAt),
		PropertyLocation:      PropertyLocationResponse(commercialPropertyForRent.Location),
		Insights:              PropertyInsightsResponse(commercialPropertyForRent.PropertyInsights),
		Media:                 ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForRent.PropertyMedia),
	}
}

func CommercialPropertyForRentWithManagerResponse(commercialPropertyForRent managerModels.CommercialRentalProperty) commercialPropertyDtos.CommercialForRentPropertyWithManagerResponseDto {
	return commercialPropertyDtos.CommercialForRentPropertyWithManagerResponseDto{
		Id:                    commercialPropertyForRent.Id,
		ManagerId:             commercialPropertyForRent.ManagerId,
		UniqueId:              commercialPropertyForRent.UniqueId,
		RentAmount:            commercialPropertyForRent.RentAmount,
		SizeNumber:            commercialPropertyForRent.SizeNumber,
		SizeDimensions:        commercialPropertyForRent.SizeDimensions,
		Status:                commercialPropertyForRent.Status,
		Type:                  commercialPropertyForRent.Type,
		YearBuilt:             commercialPropertyForRent.YearBuilt,
		Storeys:               commercialPropertyForRent.Storeys,
		HasElectricity:        commercialPropertyForRent.HasElectricity,
		IsFavorite:            false,
		HasWater:              commercialPropertyForRent.HasWater,
		NumberOfRooms:         commercialPropertyForRent.NumberOfRooms,
		MarketingStatement:    commercialPropertyForRent.MarketingStatement,
		IsFullSpace:           commercialPropertyForRent.IsFullSpace,
		TenantRequirements:    commercialPropertyForRent.TenantRequirements,
		OtherExteriorFeatures: commercialPropertyForRent.OtherExteriorFeatures,
		OtherInteriorFeatures: commercialPropertyForRent.OtherInteriorFeatures,
		Currency:              commercialPropertyForRent.Currency,
		PostedTime:            generalUtilities.GetTimePassed(commercialPropertyForRent.CreatedAt),
		PropertyLocation:      PropertyLocationResponse(commercialPropertyForRent.Location),
		Insights:              PropertyInsightsResponse(commercialPropertyForRent.PropertyInsights),
		Media:                 ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForRent.PropertyMedia),
		Manager:               managerUtilities.ManagerResponse(&commercialPropertyForRent.Manager),
	}
}

func CommercialPropertyForSaleWithManagerResponse(commercialPropertyForSale managerModels.CommercialForSaleProperty) commercialPropertyDtos.CommercialForSalePropertyWithManagerResponseDto {
	return commercialPropertyDtos.CommercialForSalePropertyWithManagerResponseDto{
		Id:                    commercialPropertyForSale.Id,
		ManagerId:             commercialPropertyForSale.ManagerId,
		UniqueId:              commercialPropertyForSale.UniqueId,
		Price:                 commercialPropertyForSale.Price,
		SizeNumber:            commercialPropertyForSale.SizeNumber,
		SizeDimensions:        commercialPropertyForSale.SizeDimensions,
		Status:                commercialPropertyForSale.Status,
		IsFavorite:            false,
		Type:                  commercialPropertyForSale.Type,
		YearBuilt:             commercialPropertyForSale.YearBuilt,
		Storeys:               commercialPropertyForSale.Storeys,
		HasElectricity:        commercialPropertyForSale.HasElectricity,
		IsNegotiable:          commercialPropertyForSale.IsNegotiable,
		HasWater:              commercialPropertyForSale.HasWater,
		MarketingStatement:    commercialPropertyForSale.MarketingStatement,
		NumberOfRooms:         commercialPropertyForSale.NumberOfRooms,
		OtherExteriorFeatures: commercialPropertyForSale.OtherExteriorFeatures,
		OtherInteriorFeatures: commercialPropertyForSale.OtherInteriorFeatures,
		Currency:              commercialPropertyForSale.Currency,
		PostedTime:            generalUtilities.GetTimePassed(commercialPropertyForSale.CreatedAt),
		PropertyLocation:      PropertyLocationResponse(commercialPropertyForSale.Location),
		Insights:              PropertyInsightsResponse(commercialPropertyForSale.PropertyInsights),
		Media:                 ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForSale.PropertyMedia),
		Manager:               managerUtilities.ManagerResponse(&commercialPropertyForSale.Manager),
	}
}

func CommercialPropertyForSaleResponse(commercialPropertyForSale managerModels.CommercialForSaleProperty) commercialPropertyDtos.CommercialForSalePropertyResponseDto {
	return commercialPropertyDtos.CommercialForSalePropertyResponseDto{
		Id:                    commercialPropertyForSale.Id,
		ManagerId:             commercialPropertyForSale.ManagerId,
		UniqueId:              commercialPropertyForSale.UniqueId,
		Price:                 commercialPropertyForSale.Price,
		SizeNumber:            commercialPropertyForSale.SizeNumber,
		SizeDimensions:        commercialPropertyForSale.SizeDimensions,
		IsFavorite:            false,
		Currency:              commercialPropertyForSale.Currency,
		Status:                commercialPropertyForSale.Status,
		Type:                  commercialPropertyForSale.Type,
		YearBuilt:             commercialPropertyForSale.YearBuilt,
		Storeys:               commercialPropertyForSale.Storeys,
		HasElectricity:        commercialPropertyForSale.HasElectricity,
		MarketingStatement:    commercialPropertyForSale.MarketingStatement,
		IsNegotiable:          commercialPropertyForSale.IsNegotiable,
		HasWater:              commercialPropertyForSale.HasWater,
		NumberOfRooms:         commercialPropertyForSale.NumberOfRooms,
		OtherExteriorFeatures: commercialPropertyForSale.OtherExteriorFeatures,
		OtherInteriorFeatures: commercialPropertyForSale.OtherInteriorFeatures,
		PostedTime:            generalUtilities.GetTimePassed(commercialPropertyForSale.CreatedAt),
		PropertyLocation:      PropertyLocationResponse(commercialPropertyForSale.Location),
		Insights:              PropertyInsightsResponse(commercialPropertyForSale.PropertyInsights),
		Media:                 ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForSale.PropertyMedia),
	}
}
