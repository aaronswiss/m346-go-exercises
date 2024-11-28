package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Grundlagen der Informatik",
		117: "Netzinfrastruktur",
		346: "Vertiefung Datenbanken",
	}
	// TODO: create a map called "modules"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])
	// TODO: delete one
	delete(modules, 104)

	// TODO: add one
	modules[431] = "Aufträge im eigenen Berufsumfeld"

	// TODO: replace one
	modules[346] = "Cloud-Lösungen planen und implementieren"

	fmt.Println(modules)
}
