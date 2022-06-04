package utils

import (
	"bufio"
	"log"
	"os"

	"github.com/Pallinder/go-randomdata"
)

// GenerateAlienNames generates random names for aliens and return slices
func GenerateAlienNames(count int) []string {
	var alienNames []string
	for i := 0; i < count; i++ {
		alienNames = append(alienNames, randomdata.SillyName())
	}
	return alienNames
}

// GenerateAlienNames and write to file
func GenerateAlienNamesAndWriteToFile(count int) {
	alienNames := GenerateAlienNames(count)
	file, err := os.OpenFile("./files/aliensNew.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()

	for i := 0; i < len(alienNames); i++ {
		datawriter.WriteString(alienNames[i] + "\n")
	}
}

// Contains checks if slice contains value
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// GeneratesCityNames takes count parameter and  generates random names for cities and return their slices
func GeneratesCityNames(count int) []string {
	// City Slice for storing generated city names
	var cityNames []string
	// while loop
	for len(cityNames) < count {
		generatedCityName := randomdata.City()
		// Check for generated city name is not in slice
		for !Contains(cityNames, generatedCityName) {
			cityNames = append(cityNames, generatedCityName)
		}
	}
	return cityNames
}
