package location

type PropertyLocationCreationDto struct {
	DisplayName string   `json:"displayName"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Surburb     string   `json:"surburb"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	Province    string   `json:"province"`
	Country     string   `json:"country"`
	CountryCode string   `json:"countryCode"`
}

type PropertyLocationUpdateAndResponseDto struct {
	Id           int      `json:"id"`
	PropertyId   int      `json:"propertyId"`
	DisplayName  string   `json:"displayName"`
	Boundingbox  []string `json:"boundingbox"`
	Lat          string   `json:"lat"`
	Lon          string   `json:"lon"`
	Surburb      string   `json:"surburb"`
	City         string   `json:"city"`
	County       string   `json:"county"`
	Province     string   `json:"province"`
	Country      string   `json:"country"`
	CountryCode  string   `json:"countryCode"`
	PropertyType string   `json:"propertyType"`
}
