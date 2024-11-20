package report

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
)

func CreatePropertyReport(report *propertyModels.PropertyReport) bool {
	db.DB.Create(report)
	return true
}

func GetAllPropertiesReports() []propertyModels.PropertyReport {
	var reports = []propertyModels.PropertyReport{}
	err := db.DB.Find(&reports)
	if err != nil {
		println(err.Error, err.Name())
	}
	return reports
}

func GetPropertyReportById(reportId string) *propertyModels.PropertyReport {
	var report propertyModels.PropertyReport
	result := db.DB.First(&report, reportId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &report
}

func GetPropertyReportsByPropertyId(propertyId string) []propertyModels.PropertyReport {
	var reports = []propertyModels.PropertyReport{}
	err := db.DB.Where("property_id= ?", propertyId).Find(&reports)
	if err != nil {
		println(err.Error, err.Name())
	}
	return reports
}

func GetPropertyReportsByManagerId(managerId string) []propertyModels.PropertyReport {
	var reports = []propertyModels.PropertyReport{}
	err := db.DB.Where("manager_id= ?", managerId).Find(&reports)
	if err != nil {
		println(err.Error, err.Name())
	}
	return reports
}

func UpdatePropertyReport(propertyReportUpdate *propertyModels.PropertyReport) bool {
	db.DB.Save(propertyReportUpdate)
	return true
}

func DeletePropertyReport(reportId string) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyReport{}, reportId)
	return true
}
