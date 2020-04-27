package api

// CountryStats struct will define the stats of a country as defined in the structure of the API
type CountryStats struct {
	Updated             int64              `json:"updated"`
	CountryName         string             `json:"country"`
	CountryInfo         CountryInformation `json:"countryInfo"`
	TotalCases          int                `json:"cases"`
	CasesToday          int                `json:"todatCases"`
	Deaths              int                `json:"deaths"`
	TodayDeaths         int                `json:"todayDeaths"`
	Recovered           int                `json:"recovered"`
	Active              int                `json:"active"`
	Critical            int                `json:"critical"`
	CasesPerOneMillion  int                `json:"casesPerOneMillion"`
	DeathsPerOneMillion int                `json:"deathsPerOneMillion"`
	Tests               int                `json:"tests"`
	TestsPerOneMillion  int                `json:"testsPerOneMillion"`
	Continent           string             `json:"continent"`
}

// CountryInformation struct will define the information about the country
type CountryInformation struct {
	ID        int     `json:"_id"`
	ISO2      int     `json:"iso2"`
	ISO3      int     `json:"iso3"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"long"`
	FlagImage string  `json:"flag"`
}
