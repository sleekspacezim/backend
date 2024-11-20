package report

type PropertyReportCreationDto struct {
	PropertyId     int    `json:"propertyId"`
	ManagerId      int    `json:"managerId"`
	ReporterUserId int    `json:"reporterUserId"`
	Report         string `json:"report"`
}

type PropertyReportUpdateAndResponseDto struct {
	Id             int    `json:"id"`
	PropertyId     int    `json:"propertyId"`
	ManagerId      int    `json:"managerId"`
	ReporterUserId int    `json:"reporterUserId"`
	Report         string `json:"report"`
}
