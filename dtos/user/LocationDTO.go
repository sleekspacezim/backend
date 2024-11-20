package dtos

type LocationDTO struct {
	Id          int      `json:"id"`
	UserId      int      `json:"userId"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Surburb     string   `json:"surburb"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	Province    string   `json:"province"`
	Country     string   `json:"country"`
	CountryCode string   `json:"countryCode"`
	DisplayName string   `json:"displayName"`
	Boundingbox []string `json:"boundingbox"`
}
