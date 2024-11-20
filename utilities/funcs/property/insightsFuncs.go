package property

import (
	propertyInsightsDtos "SleekSpace/dtos/property/insights"
	propertyModels "SleekSpace/models/property"
)

func PropertyInsightsResponse(insights propertyModels.PropertyInsights) propertyInsightsDtos.PropertyInsightsUpdateAndResponseDto {
	return propertyInsightsDtos.PropertyInsightsUpdateAndResponseDto{
		Views:             insights.Views,
		Id:                insights.Id,
		PropertyId:        insights.PropertyId,
		ContactInfoViews:  insights.ContactInfoViews,
		AddedToFavourites: insights.AddedToFavourites,
		Shared:            insights.Shared,
		PropertyType:      insights.PropertyType,
	}
}
