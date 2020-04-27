package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	allCountryStatsEndpoint = "https://disease.sh/v2/countries?yesterday=false&sort=cases"
)

// GetAllCountryStats function will return an array of country with stats of that country
// it takes no arguments and returns a slice of Country objects if successfull otherwise an error
func GetAllCountryStats() ([]CountryStats, error) {
	resp, err := http.Get(allCountryStatsEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	countries := make([]CountryStats, 0)
	json.Unmarshal(buf, &countries)
	return countries, nil
}
