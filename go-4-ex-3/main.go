package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{x1, x2}
	} else if D == 0 {
		x := -b / (2 * a)
		return []float64{x}
	} else {
		return []float64{}
	}
}

func main() {
	fmt.Println("Test der Funktion computeDiscriminant:")
	fmt.Printf("Diskriminante (a=3, b=4, c=1): %.2f\n", computeDiscriminant(3, 4, 1)) // Erwartet: 4
	fmt.Printf("Diskriminante (a=2, b=4, c=2): %.2f\n", computeDiscriminant(2, 4, 2)) // Erwartet: 0
	fmt.Printf("Diskriminante (a=3, b=4, c=2): %.2f\n", computeDiscriminant(3, 4, 2)) // Erwartet: -8

	fmt.Println("\nTest der Funktion computeQuadraticFormula:")
	fmt.Println("Testfall 1 (D > 0):")
	result1 := computeQuadraticFormula(3, 4, 1)
	fmt.Printf("Lösungen: %v\n", result1) // Erwartet: [-0.33 -1.00]

	fmt.Println("\nTestfall 2 (D = 0):")
	result2 := computeQuadraticFormula(2, 4, 2)
	fmt.Printf("Lösungen: %v\n", result2) // Erwartet: [-1.00]

	fmt.Println("\nTestfall 3 (D < 0):")
	result3 := computeQuadraticFormula(3, 4, 2)
	fmt.Printf("Lösungen: %v\n", result3) // Erwartet: []
}
