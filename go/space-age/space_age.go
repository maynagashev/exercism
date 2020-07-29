/*
Package space contains method for calculating age on other planets.

Calculate how old someone would be on:

   - Mercury: orbital period 0.2408467 Earth years
   - Venus: orbital period 0.61519726 Earth years
   - Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
   - Mars: orbital period 1.8808158 Earth years
   - Jupiter: orbital period 11.862615 Earth years
   - Saturn: orbital period 29.447498 Earth years
   - Uranus: orbital period 84.016846 Earth years
   - Neptune: orbital period 164.79132 Earth years

So if you were told someone were 1,000,000,000 seconds old, you should
be able to say that they're 31.69 Earth-years old.
*/
package space

// Planet names type
type Planet string

// EarthYearInSeconds constant value
const EarthYearInSeconds float64 = 31557600 // or 365.25 days

// EarthYearsOn other planets mapping
var EarthYearsOn = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates how old someone would be on specified planet in planet-years based on Earth age in seconds
func Age(seconds float64, planet Planet) float64 {
	planetYearInSeconds := EarthYearsOn[planet] * EarthYearInSeconds
	return seconds / planetYearInSeconds
}
