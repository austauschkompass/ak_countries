package countries

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
)

//go:embed countries.json
var countriesFile []byte

// A list of countries (with alpha2 identifiers) where some have additional subdivision
var Countries []Country

// A linear list of countries or subdivisions identified by a code like `CA` or `GB-ENG`
var CountriesByCodes []CountryOrSubdivision

type Subdivision struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	Name         string         `json:"name"`
	Alpha2       string         `json:"alpha2"`
	Subdivisions *[]Subdivision `json:"subdivisions"`
}

type CountryOrSubdivision struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Load countries on module load
func init() {
	if err := json.Unmarshal(countriesFile, &Countries); err != nil {
		// This should not happen, but if it does we should fail hard/early
		panic(err)
	}
	for _, country := range Countries {
		CountriesByCodes = append(CountriesByCodes, CountryOrSubdivision{
			Code: country.Alpha2,
			Name: country.Name,
		})

		if country.Subdivisions != nil {
			for _, subdivision := range *country.Subdivisions {
				subdivisionCode := fmt.Sprintf("%s-%s", country.Alpha2, subdivision.Code)

				CountriesByCodes = append(CountriesByCodes, CountryOrSubdivision{
					Code: subdivisionCode,
					Name: subdivision.Name,
				})
			}
		}
	}
}

// Find a country by ISO 3166-1 alpha-2 from the given list of countries
// When it cannot be found, will return nil, error!
//
// see: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
func FindCountryByAlpha2(alpha2 string) (*Country, error) {
	var country *Country
	for _, entry := range Countries {
		if entry.Alpha2 == alpha2 {
			country = &entry
			break
		}
	}

	if country == nil {
		return nil, errors.New("Failed to find Country")
	}

	return country, nil
}

// Get a name for a country by alpha2 from the given list of countries
// When given code cannot be found returns "Unbekannt"
func GetCountryNameByAlpha2(alpha2 string) string {
	country, err := FindCountryByAlpha2(alpha2)

	if err == nil {
		return country.Name
	} else {
		return "Unbekannt"
	}
}

// Find a country by code from the given list of countries
// When it cannot be found, will return nil, error!
//
// Code is either an alpha2 code (ISO 3166-1), like `CA`
// or a specific a subdivision code (ISO 3166-2), like `GB-ENG`
//
// see: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
// and: https://en.wikipedia.org/wiki/ISO_3166-2
func FindCountryByCode(code string) (*CountryOrSubdivision, error) {
	var country *CountryOrSubdivision
	for _, entry := range CountriesByCodes {
		if entry.Code == code {
			country = &entry
			break
		}
	}

	if country == nil {
		return nil, errors.New("Failed to find Country")
	}

	return country, nil
}

// Get a name for a country by code from the given list of countries
// and their subdivisions.
// When given code cannot be found returns "Unbekannt"
func GetCountryNameByCode(code string) string {
	country, err := FindCountryByCode(code)

	if err == nil {
		return country.Name
	} else {
		return "Unbekannt"
	}
}
