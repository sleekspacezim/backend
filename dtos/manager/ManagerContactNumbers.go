package dtos

import managerModels "SleekSpace/models/manager"

type ManagerContactNumbersCreationAndUpdateDTO struct {
	Contacts []managerModels.ManagerContactNumber `json:"contacts"`
}

type ManagerContactNumberDTO struct {
	Id           int    `json:"id"`
	ManagerId    int    `json:"managerId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
