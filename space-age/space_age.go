package space

var planetsMap = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Planet custom type of string
type Planet string

// float in ratio -> 1 earth year

const earthDaysPerYear = 365.25 // days in an earth year
const earthHoursPerDay = 24
const earthMinutesPerHour = 60
const earthSecondsPerMinute = 60
const earthSecondsInYear = earthSecondsPerMinute * earthMinutesPerHour * earthHoursPerDay * earthDaysPerYear

// Age to return age in years of planets based on seconds
func Age(seconds float64, planet Planet) float64 {
	earthYears := float64(seconds) / float64(earthSecondsInYear)
	planetYears := earthYears / planetsMap[planet]
	// fmt.Printf("seconds input: %v\n", seconds)
	// fmt.Printf("seconds in a year: %v\n", earthSecondsInYear)
	// fmt.Printf("earthYears: %v\n", earthYears)

	// fmt.Printf("planet: %v\n", planet)
	// fmt.Printf("planet-ratio: %v\n", planetsMap[planet])
	// fmt.Printf("%v years on planet %v\n", planetYears, planet)
	return planetYears
}
