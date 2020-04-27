package api

// Statistics struct will represent stats object of the api
// it will contain the statistics of the number of peoples either infected, dead or recovered
type Statistics struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

// Coordinates struct will represent Coordinates of a country
// it contains Latitude and Longitude of the country
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Country struct will represent the country object of the json API
type Country struct {
	Country   string      `json:"country"`
	Province  string      `json:"province"`
	UpdatedAt string      `json:"updatedAt"`
	Stats     Statistics  `json:"stats"`
	Location  Coordinates `json:"coordinates"`
}
