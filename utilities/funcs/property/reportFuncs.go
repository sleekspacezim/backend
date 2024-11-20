package property

import (
	propertyReportDtos "SleekSpace/dtos/property/report"
	propertyModels "SleekSpace/models/property"
)

func PropertyReportResponse(report propertyModels.PropertyReport) propertyReportDtos.PropertyReportUpdateAndResponseDto {
	return propertyReportDtos.PropertyReportUpdateAndResponseDto{
		Id:             report.Id,
		PropertyId:     report.PropertyId,
		ManagerId:      report.ManagerId,
		ReporterUserId: report.ReporterUserId,
		Report:         report.Report,
	}
}

func ProcessedPropertyReportsListToResponse(reports []propertyModels.PropertyReport) []propertyReportDtos.PropertyReportUpdateAndResponseDto {
	dtoList := []propertyReportDtos.PropertyReportUpdateAndResponseDto{}
	if len(reports) > 0 {
		for i := 0; i < len(reports); i++ {
			dto := propertyReportDtos.PropertyReportUpdateAndResponseDto{
				Id:             reports[i].Id,
				PropertyId:     reports[i].PropertyId,
				ManagerId:      reports[i].ManagerId,
				ReporterUserId: reports[i].ReporterUserId,
				Report:         reports[i].Report,
			}
			dtoList = append(dtoList, dto)
		}
	}
	return dtoList
}
