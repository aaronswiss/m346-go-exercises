package main

import "fmt"

func main() {
	firstName := "Aaron"
	lastName := "Ettlin"
	dayOfBirth := 18
	monthOfBirth := 04
	yearOfBirth := 2007
	numberOfSiblings := 1
	heightInMeters := 1.81
	zodiacSign := '\u2648'

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
