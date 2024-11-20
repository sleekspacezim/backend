package property

import (
	propertyStandDtos "SleekSpace/dtos/property/stand"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func PropertyStandResponse(standModel managerModels.PropertyStand) propertyStandDtos.StandResponseDTO {
	return propertyStandDtos.StandResponseDTO{
		Id:                 standModel.Id,
		ManagerId:          standModel.ManagerId,
		UniqueId:           standModel.UniqueId,
		Price:              standModel.Price,
		SizeNumber:         standModel.SizeNumber,
		SizeDimensions:     standModel.SizeDimensions,
		Status:             standModel.Status,
		IsServiced:         standModel.IsServiced,
		IsNegotiable:       standModel.IsNegotiable,
		AreaHasElectricity: standModel.AreaHasElectricity,
		MarketingStatement: standModel.MarketingStatement,
		Level:              standModel.Level,
		Currency:           standModel.Currency,
		IsFavorite:         false,
		Type:               standModel.Type,
		OtherDetails:       standModel.OtherDetails,
		PostedTime:         generalUtilities.GetTimePassed(standModel.CreatedAt),
		PropertyLocation:   PropertyLocationResponse(standModel.Location),
		Insights:           PropertyInsightsResponse(standModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(standModel.PropertyMedia),
	}
}

func PropertyStandWithManagerResponse(standModel managerModels.PropertyStand) propertyStandDtos.StandWithManagerResponseDTO {
	return propertyStandDtos.StandWithManagerResponseDTO{
		Id:                 standModel.Id,
		ManagerId:          standModel.ManagerId,
		UniqueId:           standModel.UniqueId,
		Price:              standModel.Price,
		SizeNumber:         standModel.SizeNumber,
		SizeDimensions:     standModel.SizeDimensions,
		Status:             standModel.Status,
		IsServiced:         standModel.IsServiced,
		IsNegotiable:       standModel.IsNegotiable,
		AreaHasElectricity: standModel.AreaHasElectricity,
		Level:              standModel.Level,
		IsFavorite:         false,
		Type:               standModel.Type,
		MarketingStatement: standModel.MarketingStatement,
		OtherDetails:       standModel.OtherDetails,
		Currency:           standModel.Currency,
		PostedTime:         generalUtilities.GetTimePassed(standModel.CreatedAt),
		PropertyLocation:   PropertyLocationResponse(standModel.Location),
		Insights:           PropertyInsightsResponse(standModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(standModel.PropertyMedia),
		Manager:            managerUtilities.ManagerResponse(&standModel.Manager),
	}
}
