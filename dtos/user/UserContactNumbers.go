package dtos

type ContactNumberDTO struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
