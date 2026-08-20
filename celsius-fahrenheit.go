package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	celsius := 37.0

	fahrenheit := celsiusToFahrenheit(celsius)

	fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)

	fahrenheit = 98.6

	celsius = fahrenheitToCelsius(fahrenheit)

	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
}
