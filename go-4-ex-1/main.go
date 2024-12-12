package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints float64, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 {
		return 0, errors.New("Punktzahlen dürfen nicht negativ sein, und maximale Punktzahl muss positiv sein")
	}
	if gotPoints > maxPoints {
		return 0, errors.New("erreichte Punktzahl kann nicht größer als die maximale Punktzahl sein")
	}

	grade := (gotPoints/maxPoints)*5 + 1
	return grade, nil
}

func main() {
	tests := []struct {
		gotPoints float64
		maxPoints float64
	}{
		{17.5, 28.0},  // 4.125
		{24.0, 30.0},  // 5.0
		{-5.0, 30.0},  // 2.25
		{40.0, 30.0},  // -
		{10.0, -20.0}, // -
		{10.0, 40.0},  //2.25
	}

	for _, test := range tests {
		grade, err := computeGrade(test.gotPoints, test.maxPoints)
		if err != nil {
			fmt.Printf("Fehler für Eingabe (gotPoints=%.2f, maxPoints=%.2f): %s\n", test.gotPoints, test.maxPoints, err)
		} else {
			fmt.Printf("Note für (gotPoints=%.2f, maxPoints=%.2f): %.3f\n", test.gotPoints, test.maxPoints, grade)
		}
	}
}
