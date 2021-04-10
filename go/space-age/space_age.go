// Package space calculates a person's age in years on different planets.
package space

type Planet string

// Dafauts earth periods
const Earth = 31557600.0

// Define map with planet name and peridos
var PlanetSeconds = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132,
}

// Age accepts age in seconds and a planet name and returns a person's age on that planet
func Age(secs float64, planet Planet) float64 {
	return secs / PlanetSeconds[planet]
}
