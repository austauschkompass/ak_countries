const fs = require('fs')
const {
  main: {
    de: {
      localeDisplayNames: {
        territories
      }
    }
  }
} = require('cldr-localenames-full/main/de/territories.json')
const fetch = require('node-fetch')

const countries = []
const seenAlpha2Codes = new Set()

fetch('https://datahub.io/core/country-list/r/data.json')
  .then((data) => data.json())
  .then((json) => extractLocalizedNames(json))
  .catch((err) => console.error(err))

function extractLocalizedNames(alpha2List) {
  for (const country of alpha2List) {
    const { Code: alpha2, Name: englishName } = country
    let name = territories[alpha2]
    const shortName = territories[`${alpha2}-alt-short`]
    // prefer shortname if available
    if (shortName) {
      name = shortName
    }
    countries.push({
      alpha2,
      name
    })
  }
  // Add in Subdivisions for United Kingdom
  countries.find(({ alpha2 }) => alpha2 === 'GB').subdivisions = [
    {
      code: 'ENG',
      name: 'England'
    },
    {
      code: 'NIR',
      name: 'Nordirland'
    },
    {
      code: 'SCT',
      name: 'Schottland'
    },
    {
      code: 'WLS',
      name: 'Wales'
    }
  ]
  
  fs.writeFileSync('countries.json', prettyJSON(sortCountries(countries)))
}

function prettyJSON(object) {
  return JSON.stringify(object, null, 2)
}

function sortCountries(countries) {
  return countries.sort((a,b) => a.name.localeCompare(b.name))
}
