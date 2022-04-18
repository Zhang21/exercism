// Package weather is xxx.
package weather

// CurrentCondition is xxx.
var CurrentCondition string

// CurrentLocation is xxx.
var CurrentLocation string

// Forecast func is xxx.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
