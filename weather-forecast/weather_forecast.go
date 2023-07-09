// Package weather / Package weather provides current weather condition for a given city in Goblinocus.
package weather

// CurrentCondition is the current weather condition for a given city in Goblinocus.
var CurrentCondition string

// CurrentLocation is the current weather condition for a given city in Goblinocus.
var CurrentLocation string

// Forecast / Forecast returns the current weather condition for a given city in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
