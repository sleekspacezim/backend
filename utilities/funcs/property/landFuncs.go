package property

import (
	landDtos "SleekSpace/dtos/property/land"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func LandPropertyResponse(landPropertyModel managerModels.LandForSaleProperty) landDtos.LandForSalePropertyResponseDto {
	return landDtos.LandForSalePropertyResponseDto{
		Id:                 landPropertyModel.Id,
		ManagerId:          landPropertyModel.ManagerId,
		UniqueId:           landPropertyModel.UniqueId,
		Price:              landPropertyModel.Price,
		SizeNumber:         landPropertyModel.SizeNumber,
		SizeDimensions:     landPropertyModel.SizeDimensions,
		Status:             landPropertyModel.Status,
		IsFavorite:         false,
		Type:               landPropertyModel.Type,
		AreaHasElectricity: landPropertyModel.AreaHasElectricity,
		HasWater:           landPropertyModel.HasWater,
		IsNegotiable:       landPropertyModel.IsNegotiable,
		OtherDetails:       landPropertyModel.OtherDetails,
		Currency:           landPropertyModel.Currency,
		MarketingStatement: landPropertyModel.MarketingStatement,
		PostedTime:         generalUtilities.GetTimePassed(landPropertyModel.CreatedAt),
		PropertyLocation:   PropertyLocationResponse(landPropertyModel.Location),
		Insights:           PropertyInsightsResponse(landPropertyModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(landPropertyModel.PropertyMedia),
	}
}

func LandPropertyWithManagerResponse(landPropertyModel managerModels.LandForSaleProperty) landDtos.LandForSalePropertyWithManagerResponseDto {
	return landDtos.LandForSalePropertyWithManagerResponseDto{
		Id:                 landPropertyModel.Id,
		ManagerId:          landPropertyModel.ManagerId,
		UniqueId:           landPropertyModel.UniqueId,
		Price:              landPropertyModel.Price,
		SizeNumber:         landPropertyModel.SizeNumber,
		SizeDimensions:     landPropertyModel.SizeDimensions,
		Status:             landPropertyModel.Status,
		Type:               landPropertyModel.Type,
		Currency:           landPropertyModel.Currency,
		IsFavorite:         false,
		MarketingStatement: landPropertyModel.MarketingStatement,
		AreaHasElectricity: landPropertyModel.AreaHasElectricity,
		HasWater:           landPropertyModel.HasWater,
		IsNegotiable:       landPropertyModel.IsNegotiable,
		OtherDetails:       landPropertyModel.OtherDetails,
		PostedTime:         generalUtilities.GetTimePassed(landPropertyModel.CreatedAt),
		PropertyLocation:   PropertyLocationResponse(landPropertyModel.Location),
		Insights:           PropertyInsightsResponse(landPropertyModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(landPropertyModel.PropertyMedia),
		Manager:            managerUtilities.ManagerResponse(&landPropertyModel.Manager),
	}
}
