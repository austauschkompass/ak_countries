# AustauschKompass Countries

Single source of truth for a localized list of country names. Consumable as an [npm package](#javascript-npm-package) and [go module](#go-module).

Contents are localized to german.

When available, we chose the short form for countries, e.g. "USA"
instead of "Vereinigte Staaten von Amerika", "UK" instead of
"Vereinigtes Königreich".

## JavaScript (npm package)

### Installation

to install in your project run:

```
yarn add git+ssh://git@github.com:austauschkompass/ak_countries.git
```

__NOTE__: We directly load the raw JSON within `index.js` so if you
bundle your app for the browser (e.g. via `webpack`/`rollup`) you have
to ensure JSON is loaded and embedded (e.g. via `json-loader` for
`webpack`) correctly. Otherwise the imported lists will be
undefined!

### Usage

This library exports various country lists for different usage
scenarios.

All lists contain objects that have a and are
(locale aware) sorted by `name`.

### `akCountriesWithSubdivisions`

Use this list if you want to allow users more fine-grained selection
of countries/subdivisions.

This list contains all countries and their subdivisions as objects containing a `code` and a `name`.
The `code` can be used as an identifier and is either:

1. a two-letter country code, e.g. `US` (ISO 3166-1)
2. a subdivision code, e.g. `GB-NIR` (ISO 3166-2) composed of:
   1. a two-letter cc, e.g. `GB`
   2. the dash `-`)
   2. an up to three letter alpha-numeric identifier, e.g. `NIR`

Be aware that all consuming/producing interfaces (Apps/Databases/APIs)
have to be aware of the possible values in `code`.

__Schema:__

```json5
[
  {
    "code": "AF"
    "name": "Afghanistan"
  },
  // ...
  {
    "code": "GB-ENG"
    "name": "England (UK)"
  },
  // ...
  {
    "code": "GB"
    "name": "UK"
  },
  // ...
]
```

__Example:__

```js
import { akCountriesWithSubdivisions } from 'ak_countries'

// example finder function for subdivision aware search
const findSubdivisionName = (query) => akCountriesWithSubdivisions.find(({ code }) => query === code)

const countryName = findSubdivisionName('GB') // "UK"
const countryName = findSubdivisionName('MT') // "Malta"
// and
const countryName = findSubdivisionName('GB-WLS') // "Wales"
```

### `akCountries`

Use this if you do not care about or cannot use the compound ISO 3166-2
subdivision codes, otherwise use `akCountriesWithSubdivisions`.

This list contains only the 249 Countries as listed by ISO 3166-1.
Subdivisions are added hierarchically below the corresponding entry.

__Schema:__

```json5
[
  {
    "alpha2": "AF"
    "name": "Afghanistan"
  },
  // ...
  {
    "alpha2": "GB",
    "name": "UK",
    "subdivisions": [
      {
        "code": "ENG"
        "name": "England"
      },
      // ...
    ]
  }
  // ...
]
```

__Example:__

```js
import { akCountries } from 'ak_countries'

// example finder function (ignoring subdivisions)
const findCountryName = (query) => akCountries.find(({ alpha2 }) => query.slice(0,2) === alpha2)

const countryName = findCountryName('GB') // "UK"
const countryName = findCountryName('MT') // "Malta"
// but
const countryName = findCountryName('GB-WLS') // "UK"
```

## Go (module)

__TODO__ Handle subdivision lookup

Example usage within your go project:

```go
import (
  "log"
  "github.com/austauschkompass/ak_countries"
)

func main() {
  countries, err := ak_countries.LoadCountries()
  
  if err != nil {
    log.Fatalf("Failed loading Countries: %v", err)
  }
  
  log.Printf("Wie, du warst noch niemals in %s?", ak_countries.GetCountryNameByCode(countries, "CA"))
}

```

Run `go mod tidy` once.


### `GetCountryNameByCode`

This function will always return a string and fallback to "Unbekannt" if the given country cannot be found.

### `FindCountryByCode`

This function returns an error if the given country cannot be found.

Example usage (assuming `countries` was already loaded and checked via `LoadCountries`):

```
country, err := ak_countries.FindCountryByCode(countries, "MT")

if err == nil {
  log.Printf("Mit dem Rad nach %s (%s)?", country.Name, country.Alpha2)
}
```

## How to update data and where does it come from?

The list of official two-letter country codes (ISO 3166-1) is
downloaded from [datahub](https://datahub.io/core/country-list) and
the localized names taken from the [CLDR
project](https://github.com/unicode-org/cldr-json).

You can update and merge in newer translations (in case territories
change their names etc.) by running:

```
yarn install
yarn run refresh-localized-names
```

This will rewrite `countries.json` so be sure to check in changes one
by one, in case manual corrections were overwritten.

Once this is done, commit the results and push changes upstream, to be able to
update your Apps/Components to use the new translations via:

```
yarn upgrade ak_countries
```

__NOTE__: We merge in explicit names for subdivisions of the United Kingdom.
