package insights

type PropertyInsightsCreationDto struct {
	Views             int    `json:"views"`
	AddedToFavourites int    `json:"addedToFavourites"`
	Shared            int    `json:"shared"`
	ContactInfoViews  int    `json:"contactInfoViews"`
	EmailAttempts     int    `json:"emailAttempts"`
	CallAttempts      int    `json:"callAttempts"`
	WhatsAppAttempts  int    `json:"whatsAppAttempts"`
	PropertyType      string `json:"propertyType"`
}

type PropertyInsightsUpdateAndResponseDto struct {
	Id                int    `json:"id"`
	PropertyId        int    `json:"propertyId"`
	Views             int    `json:"views"`
	AddedToFavourites int    `json:"addedToFavourites"`
	Shared            int    `json:"shared"`
	ContactInfoViews  int    `json:"contactInfoViews"`
	EmailAttempts     int    `json:"emailAttempts"`
	CallAttempts      int    `json:"callAttempts"`
	WhatsAppAttempts  int    `json:"whatsAppAttempts"`
	PropertyType      string `json:"propertyType"`
}
