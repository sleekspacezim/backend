package insights

import (
	"net/http"

	insightsDtos "SleekSpace/dtos/property/insights"
	propertyModels "SleekSpace/models/property"
	insightsRepo "SleekSpace/repositories/property/insights"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetPropertyInsightsById(c *gin.Context) {
	insights := insightsRepo.GetPropertyInsightsById(c.Param("id"))
	if insights == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property insights information does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyInsightsResponse(*insights)})
}

func GetPropertyInsightsByPropertyId(c *gin.Context) {
	insights := insightsRepo.GetPropertyInsightsByPropertyId(c.Param("id"))
	if insights == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property insights information does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyInsightsResponse(*insights)})
}

func UpdatePropertyInsights(c *gin.Context) {
	var insightsUpdateDetails insightsDtos.PropertyInsightsUpdateAndResponseDto
	validateModelFields := validator.New()
	c.BindJSON(&insightsUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(insightsUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	insightsUpdate := propertyModels.PropertyInsights{
		Id:                insightsUpdateDetails.Id,
		PropertyId:        insightsUpdateDetails.PropertyId,
		PropertyType:      insightsUpdateDetails.PropertyType,
		Views:             insightsUpdateDetails.Views,
		ContactInfoViews:  insightsUpdateDetails.ContactInfoViews,
		AddedToFavourites: insightsUpdateDetails.AddedToFavourites,
		Shared:            insightsUpdateDetails.Shared,
	}

	isInsightsUpdated := insightsRepo.UpdatePropertyInsights(&insightsUpdate)
	if !isInsightsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property insights"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": true})
}
