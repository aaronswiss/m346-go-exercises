package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var cozy Celsius = 23.0
	var cold Fahrenheit = 15.3

	fmt.Println("Test der Methoden:")
	fahrenheitFromCelsius := cozy.ConvertToFahrenheit()
	fmt.Printf("%.2f °C = %.2f °F\n", cozy, fahrenheitFromCelsius) // 23.00 °C = 73.40 °F

	celsiusFromFahrenheit := cold.ConvertToCelsius()
	fmt.Printf("%.2f °F = %.2f °C\n", cold, celsiusFromFahrenheit) // 15.30 °F = -9.28 °C

	fmt.Println("\nRückumwandlung zum Test der Genauigkeit:")
	backToCelsius := fahrenheitFromCelsius.ConvertToCelsius()
	fmt.Printf("%.2f °F zurück in Celsius: %.2f °C\n", fahrenheitFromCelsius, backToCelsius) // 73.40 °F zurück in Celsius: 23.00 °C

	backToFahrenheit := celsiusFromFahrenheit.ConvertToFahrenheit()
	fmt.Printf("%.2f °C zurück in Fahrenheit: %.2f °F\n", celsiusFromFahrenheit, backToFahrenheit) // -9.28 °C zurück in Fahrenheit: 15.30 °F
}

// Variante 2 ist übersichtlicher, da sie die Objektorientierung nutzt.
// Durch Methodenaufrufe an den Typen selbst (Celsius oder Fahrenheit) wird der Code intuitiver und einfacher zu lesen,
// da klar ist, welche Umwandlung zu welchem Typ gehört.
