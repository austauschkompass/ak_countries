var countries = require ('./countries.json')

/*
   Convert hierarchical list of countries and their subdivisions into
   flat array, e.g. in addition to United Kingdom (GB) it will contain
   four entries for the separate subdivisions (England, Scotland…).

   List contains objects of the form: [ { name: "England", code:
   "GB-ENG" }, { name: "Schweiz", code: "CH" } ]

   Useful for e.g. linear UI Elements like <select/>
*/
function flattenSubdivisions(countries) {
  const withSubdivisions = []
  for (let current of countries) {
    if (current.subdivisions) {
      for (let subdivision of current.subdivisions) {
        withSubdivisions.push({
          code: `${current.alpha2}-${subdivision.code}`,
          name: `${subdivision.name} (${current.name})`
        })
      }
    }
    withSubdivisions.push({
      code: current.alpha2,
      name: current.name
    })
  }
  return withSubdivisions
}

function localeSortedByName (a, b) {
  return a.name.localeCompare(b.name)
}

exports.akCountriesWithSubdivisions = flattenSubdivisions(countries).sort(localeSortedByName)
exports.akCountries = countries.sort(localeSortedByName)
