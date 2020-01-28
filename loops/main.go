package main

import (
	"log"
)

func main() {
	cities := []string{"Santa Clara", "Palo Alto", "San Francisco", "Mountain View"}
	log.Println("Here are some cities:")
	for i, city := range cities {
		log.Printf("\t%d: %s\n", i, city)
	}

	moreCities := []string{"Seattle", "New York", "Orlando", "Portland"}

	log.Println("Here are some more cities:")
	allCities := append(cities, moreCities...)
	for i, city := range allCities {
		log.Printf("\t%d: %s\n", i, city)
	}
}
