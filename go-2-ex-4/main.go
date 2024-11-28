package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)

	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)

	type Class struct {
		Name     string
		Students []Student
	}

	Class1 := []Student{
		{FirstName: "Eneas", LastName: "Inanger"},
		{FirstName: "Phil", LastName: "Hägeli"},
		{FirstName: "Tim", LastName: "Hüsler"},
	}

	Class2 := []Student{
		{FirstName: "Lukas", LastName: "Kunz"},
		{FirstName: "Lukas", LastName: "Kunz"},
		{FirstName: "Lukas", LastName: "Kunz"},
	}

	classA := Class{Name: "Class A", Students: Class1}
	classB := Class{Name: "Class B", Students: Class2}

	// TODO: declare a map of modules being attended by multiple classes

	modules := map[int][]Class{
		346: {classA, classB},
		567: {classB},
		789: {classA},
	}

	// TODO: output everything using fmt.Println()

	fmt.Println("Modules and their classes:")
	for moduleID, classes := range modules {
		fmt.Printf("Module ID: %d\n", moduleID)
		for _, class := range classes {
			fmt.Printf("  Class: %s\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    Student: %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
