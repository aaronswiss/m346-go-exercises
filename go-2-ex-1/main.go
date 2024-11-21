package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  uint16
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			FirstName: "Aaron",
			LastName:  "Ettlin",
		},
		BirthDate: BirthDate{
			DayOfBirth:   18,
			MonthOfBirth: 04,
			YearOfBirth:  2007,
		},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u2648', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
}
