package countries

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestLoadCountries(t *testing.T) {
	countries, err := LoadCountries()

	if err != nil {
		t.Fatalf("Expected loading countries to succeed, but if failed with: %v", err)
	}

	if len(countries) == 0 {
		t.Fatalf("Expected some countries, but the list is empty")
	}

	country := countries[0]

	if country.Alpha2 != "AF" {
		t.Fatalf("Expected AF as first country code, but got: %v", country.Alpha2)
	}

	if country.Name != "Afghanistan" {
		t.Fatalf("Expected Afghanistan as first country name, but got: %v", country.Name)
	}

	var countryWithSubdivisions *Country

	for _, c := range countries {
		if c.Subdivisions != nil {
			countryWithSubdivisions = &c
			break
		}
	}

	if countryWithSubdivisions == nil {
		t.Fatalf("Expected at least on country with subdivisions")
	}

	subdivisions := *countryWithSubdivisions.Subdivisions

	if len(subdivisions) == 0 {
		t.Fatalf("Expected at least one subdivision")
	}

	if subdivisions[0].Code != "ENG" {
		t.Fatalf("Expected first subdivision to be ENG, but was: %v", subdivisions[0].Code)
	}

	if subdivisions[0].Name != "England" {
		t.Fatalf("Expected first subdivision name to be England, but was: %v", subdivisions[0].Name)
	}
}

func TestFindCountryByCode(t *testing.T) {
	countries, err := LoadCountries()

	if err != nil {
		t.Fatalf("failed to load countries")
	}

	country, err := FindCountryByCode(countries, "CA")

	if err != nil {
		t.Fatalf("Failed finding country with code CA: %v", err)
	}

	if country.Name != "Kanada" {
		t.Fatalf("Expected Country with name 'Kanada', but got: %v", country.Name)
	}
}

func TestFindCountryByCodeFail(t *testing.T) {
	countries, err := LoadCountries()

	if err != nil {
		t.Fatalf("failed to load countries")
	}

	_, err = FindCountryByCode(countries, "DEFINITELY_NOT_THERE")

	if err == nil {
		t.Fatalf("Expected finding country to fail, but it did not")
	}
}

func TestGetCountryNameByCode(t *testing.T) {
	countries, err := LoadCountries()

	if err != nil {
		t.Fatalf("failed to load countries")
	}

	name := GetCountryNameByCode(countries, "CA")

	if name != "Kanada" {
		t.Fatalf("Expected Country name to be 'Kanada' but got: %v", name)
	}
}
