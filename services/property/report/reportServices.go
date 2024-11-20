package report

import (
	"net/http"

	reportDtos "SleekSpace/dtos/property/report"
	propertyModels "SleekSpace/models/property"
	reportRepo "SleekSpace/repositories/property/report"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePropertyReportById(c *gin.Context) {
	var propertyReportCreationDetails reportDtos.PropertyReportCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&propertyReportCreationDetails)

	modelFieldsValidationError := validateModelFields.Struct(propertyReportCreationDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	report := propertyModels.PropertyReport{
		PropertyId:     propertyReportCreationDetails.PropertyId,
		ManagerId:      propertyReportCreationDetails.ManagerId,
		ReporterUserId: propertyReportCreationDetails.ReporterUserId,
		Report:         propertyReportCreationDetails.Report,
	}

	isReportUpdated := reportRepo.UpdatePropertyReport(&report)
	if !isReportUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create your report."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": true})
}

func GetPropertyReportById(c *gin.Context) {
	report := reportRepo.GetPropertyReportById(c.Param("id"))
	if report == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this report does not exist."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyReportResponse(*report)})
}

func GetPropertyReportsByPropertyId(c *gin.Context) {
	reports := reportRepo.GetPropertyReportsByPropertyId(c.Param("id"))
	if reports == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not have any reports."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ProcessedPropertyReportsListToResponse(reports)})
}

func GetPropertyReportsByManagerId(c *gin.Context) {
	reports := reportRepo.GetPropertyReportsByManagerId(c.Param("id"))
	if reports == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not have any reports."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ProcessedPropertyReportsListToResponse(reports)})
}

func UpdatePropertyReportById(c *gin.Context) {
	var propertyReportUpdateDetails reportDtos.PropertyReportUpdateAndResponseDto
	validateModelFields := validator.New()
	c.BindJSON(&propertyReportUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(propertyReportUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	reportUpdate := propertyModels.PropertyReport{
		Id:             propertyReportUpdateDetails.Id,
		PropertyId:     propertyReportUpdateDetails.PropertyId,
		ManagerId:      propertyReportUpdateDetails.ManagerId,
		ReporterUserId: propertyReportUpdateDetails.ReporterUserId,
		Report:         propertyReportUpdateDetails.Report,
	}

	isReportUpdated := reportRepo.UpdatePropertyReport(&reportUpdate)
	if !isReportUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your report."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": true})
}

func DeletePropertyReport(c *gin.Context) {
	report := reportRepo.GetPropertyReportById(c.Param("id"))
	if report == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this report does not exist."})
		return
	}
	isReportDeleted := reportRepo.DeletePropertyReport(c.Param("id"))
	if !isReportDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this report does not exist."})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": true})
	}
}
