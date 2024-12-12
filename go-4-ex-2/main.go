package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	fmt.Println("Testfälle für die Berechnung der Hypotenuse (computeHypotenuse):")

	a1, b1 := 3.0, 4.0
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", a1, b1, computeHypotenuse(a1, b1)) // 5.0

	a2, b2 := 5.0, 12.0
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", a2, b2, computeHypotenuse(a2, b2)) // 13.0

	a3, b3 := 8.0, 15.0
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", a3, b3, computeHypotenuse(a3, b3)) // 17.0

	a4, b4 := 7.0, 24.0
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", a4, b4, computeHypotenuse(a4, b4)) // 25.0

	fmt.Println("\nTestfälle für die Berechnung der Hypotenuse (ShortSides.Hypotenuse):")

	side1 := ShortSides{a: 3, b: 4}
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", side1.a, side1.b, side1.Hypotenuse()) // 5.0

	side2 := ShortSides{a: 5, b: 12}
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", side2.a, side2.b, side2.Hypotenuse()) // 13.0

	side3 := ShortSides{a: 8, b: 15}
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", side3.a, side3.b, side3.Hypotenuse()) // 17.0

	side4 := ShortSides{a: 7, b: 24}
	fmt.Printf("a = %.1f, b = %.1f, c = %.2f\n", side4.a, side4.b, side4.Hypotenuse()) // 25.0
}
