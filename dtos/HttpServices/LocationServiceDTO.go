package dtos

type LocationQueryDTO struct {
	LocationName string `json:"locationName"`
}

type ReverseGeoCodingDTO struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type IAddress struct {
	City         string `json:"city"`
	Country      string `json:"country"`
	Country_code string `json:"country_code"`
	County       string `json:"county"`
	State        string `json:"state"`
	Suburb       string `json:"surburb"`
}

type ISearchLocation struct {
	Address         IAddress `json:"address"`
	Boundingbox     []string `json:"boundingbox"`
	Class           string   `json:"class"`
	Display_address string   `json:"display_address"`
	Display_name    string   `json:"display_name"`
	Display_place   string   `json:"display_place"`
	Lat             string   `json:"lat"`
	Licence         string   `json:"license"`
	Lon             string   `json:"lon"`
	Osm_id          string   `json:"osm_id"`
	Osm_type        string   `json:"osm_type"`
	Place_id        string   `json:"place_id"`
	Type            string   `json:"type"`
}

type IReverseLocation struct {
	Address      IAddress `json:"address"`
	Boundingbox  []string `json:"boundingbox"`
	Display_name string   `json:"display_name"`
	Lat          string   `json:"lat"`
	Licence      string   `json:"license"`
	Lon          string   `json:"lon"`
	Osm_id       string   `json:"osm_id"`
	Osm_type     string   `json:"osm_type"`
	Place_id     string   `json:"place_id"`
}
