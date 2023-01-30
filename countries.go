package countries

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed countries.json
var countriesFile []byte

var countries []Country

type Subdivision struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	Name         string         `json:"name"`
	Alpha2       string         `json:"alpha2"`
	Subdivisions *[]Subdivision `json:"subdivisions"`
}

// Load countries once, then return global variable directly
func LoadCountries() ([]Country, error) {
	if len(countries) == 0 {
		if err := json.Unmarshal(countriesFile, &countries); err != nil {
			return countries, err
		}
	}

	return countries, nil
}

// Find a country by code from the given list of countries
// When the given code cannot be found, will return nil, error!
func FindCountryByCode(countries []Country, code string) (*Country, error) {
	var country *Country

	for _, entry := range countries {
		if entry.Alpha2 == code {
			country = &entry
			break
		}
	}

	if country == nil {
		return nil, errors.New("Failed to find Country")
	} else {
		return country, nil
	}
}

// Get a name for a country by code from the given list of countries
// When given code cannot be found returns "Unbekannt"
func GetCountryNameByCode(countries []Country, code string) string {
	country, err := FindCountryByCode(countries, code)

	if err == nil {
		return country.Name
	} else {
		return "Unbekannt"
	}
}
