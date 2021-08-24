# ak_countries

List of Countries with localized names. Consumable as an npm package.

Its contents are localized to german and queries against the data is expected to use
ISO 3166 two-letter country codes (e.g. 'US').

When available, we chose the short form for countries, e.g. "USA"
instead of "United States of America".

## Installation

to install in your project run:

```
yarn add git+ssh://git@github.com:austauschkompass/ak_countries.git
```

__NOTE__: We directly load the raw JSON within `index.js` so if you
bundle your app for the browser (e.g. via `webpack`/`rollup`) you have
to ensure JSON is correctly loaded and embedded (e.g. use
`json-loader` for `webpack`). Otherwise the imported `akCountries`
will be undefined!

## Usage

```js
import { akCountries } from 'ak_countries'

// example finder function for subdivision aware search
const findCountryName = (query) => {
  alpha2Query, subdivisionQuery = query.split('-')
  const country = akCountries.find(({ alpha2 }) => alpha2Query === alpha2)

  if (subdivisionQuery) {
    const subdivision = country.subdivisions.find(({ code }) => code === subdivisionQuery)
    return subdivision.name
  }
    return country.name
  }
}

const countryName = findCountryName('GB-WLS') // "Wales"
const countryName = findCountryName('GB') // "Vereinigtes Königreich"
const countryName = findCountryName('MT') // "Malta"
```

## How to update data and where does it come from?

The list of official Two-letter country codes (ISO 3166-1) is
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
