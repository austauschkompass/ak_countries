package countries

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed countries.json
var countriesFile []byte

var Countries []Country

type Subdivision struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	Name         string         `json:"name"`
	Alpha2       string         `json:"alpha2"`
	Subdivisions *[]Subdivision `json:"subdivisions"`
}

// Load countries on module load
func init() {
	if err := json.Unmarshal(countriesFile, &Countries); err != nil {
		// This should not happen, but if it does we should fail hard/early
		panic(err)
	}
}

// Find a country by code from the given list of countries
// When the given code cannot be found, will return nil, error!
func FindCountryByCode(code string) (*Country, error) {
	var country *Country

	for _, entry := range Countries {
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
func GetCountryNameByCode(code string) string {
	country, err := FindCountryByCode(code)

	if err == nil {
		return country.Name
	} else {
		return "Unbekannt"
	}
}
