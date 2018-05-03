package space

// Planet is an alias for a string
type Planet string

const (
	// Mercury is the first Planet in our solar system
	Mercury     Planet = "Mercury"
	mercuryYear        = 0.2408467 * earthYear

	// Venus is the second Planet in our solar system
	Venus     Planet = "Venus"
	venusYear        = 0.61519726 * earthYear

	// Earth is the third Planet in our solar system
	Earth     Planet  = "Earth"
	earthYear float64 = 31557600

	// Mars is the fourth Planet in our solar system
	Mars     Planet = "Mars"
	marsYear        = 1.8808158 * earthYear

	// Jupiter is the fifth Planet in our solar system
	Jupiter     Planet = "Jupiter"
	jupiterYear        = 11.862615 * earthYear

	// Saturn is the sixth Planet in our solar system
	Saturn     Planet = "Saturn"
	saturnYear        = 29.447498 * earthYear

	// Uranus is the seventh Planet in our solar system
	Uranus     Planet = "Uranus"
	uranusYear        = 84.016846 * earthYear

	// Neptune is the eigth Planet in our solar system
	Neptune     Planet = "Neptune"
	neptuneYear        = 164.79132 * earthYear
)

var orbits = map[Planet]float64{
	Mercury: mercuryYear,
	Venus:   venusYear,
	Earth:   earthYear,
	Mars:    marsYear,
	Jupiter: jupiterYear,
	Saturn:  saturnYear,
	Uranus:  uranusYear,
	Neptune: neptuneYear,
}

// Age gives the age of someone in Earth years
func Age(age float64, planet Planet) float64 {
	return age / orbits[planet]
}
