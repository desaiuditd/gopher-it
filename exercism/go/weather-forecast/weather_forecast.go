// Package weather provide a function and properties to describe current weather for a city.
package weather

// CurrentCondition shows current weather condition for the current city.
// Its value is set from `Forecast` method.
var CurrentCondition string

// CurrentLocation holds the name of current city, we are interested to know the weather in.
// Its value is set from `Forecast` method.
var CurrentLocation string

// Forecast returns a string telling you about the current city and its current weather condition.
// You can use this string in your UI to inform the user about the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
