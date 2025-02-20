package insights

import (
	"net/http"

	insightsDtos "SleekSpace/dtos/property/insights"
	propertyModels "SleekSpace/models/property"
	insightsRepo "SleekSpace/repositories/property/insights"
	generalUtilities "SleekSpace/utilities/funcs/general"
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
		WhatsAppAttempts:  insightsUpdateDetails.WhatsAppAttempts,
		EmailAttempts:     insightsUpdateDetails.EmailAttempts,
		CallAttempts:      insightsUpdateDetails.CallAttempts,
	}

	isInsightsUpdated := insightsRepo.UpdatePropertyInsights(&insightsUpdate)
	if !isInsightsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property insights"})
		return
	}
	insights := insightsRepo.GetPropertyInsightsById(
		generalUtilities.ConvertIntToString(insightsUpdateDetails.Id),
	)
	if insights == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property insights information does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyInsightsResponse(*insights)})
}

func IncreamentInsightsProperties(c *gin.Context) {
	var insightsUpdateDetails struct {
		InsightProperty string `json:"insightProperty"`
	}
	validateModelFields := validator.New()
	c.BindJSON(&insightsUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(insightsUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	insights := insightsRepo.GetPropertyInsightsByPropertyId(
		c.Param("id"),
	)
	if insights == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property insights information does not exist"})
		return
	}

	views := insights.Views
	emailAttempts := insights.EmailAttempts
	callAttempts := insights.CallAttempts
	whatsAppAttempts := insights.WhatsAppAttempts
	shared := insights.Shared
	addedToFavourites := insights.AddedToFavourites
	contactInfoViews := insights.ContactInfoViews

	if insightsUpdateDetails.InsightProperty == "views" {
		views = insights.Views + 1
	}
	if insightsUpdateDetails.InsightProperty == "emailAttempts" {
		emailAttempts = insights.EmailAttempts + 1
	}
	if insightsUpdateDetails.InsightProperty == "callAttempts" {
		callAttempts = insights.CallAttempts + 1
	}
	if insightsUpdateDetails.InsightProperty == "whatsAppAttempts" {
		whatsAppAttempts = insights.WhatsAppAttempts + 1
	}
	if insightsUpdateDetails.InsightProperty == "shared" {
		shared = insights.Shared + 1
	}
	if insightsUpdateDetails.InsightProperty == "addedToFavourites" {
		println("added", insights.AddedToFavourites+1)
		addedToFavourites = insights.AddedToFavourites + 1
	}
	if insightsUpdateDetails.InsightProperty == "removedFromFavourites" {
		println("added", insights.AddedToFavourites-1)
		addedToFavourites = insights.AddedToFavourites - 1
	}
	if insightsUpdateDetails.InsightProperty == "contactInfoViews" {
		contactInfoViews = insights.ContactInfoViews + 1
	}

	insightsUpdate := propertyModels.PropertyInsights{
		Id:                insights.Id,
		PropertyId:        insights.PropertyId,
		PropertyType:      insights.PropertyType,
		Views:             views,
		ContactInfoViews:  contactInfoViews,
		AddedToFavourites: addedToFavourites,
		Shared:            shared,
		WhatsAppAttempts:  whatsAppAttempts,
		EmailAttempts:     emailAttempts,
		CallAttempts:      callAttempts,
	}

	isInsightsUpdated := insightsRepo.UpdatePropertyInsights(&insightsUpdate)
	if !isInsightsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property insights"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": true})
}
