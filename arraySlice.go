package main

import "fmt"

func main() {
	// array
	// fixed size
	names := [5]string{"Shihab", "Shana", "Anwar", "Shifa", "Althaf"}
	fmt.Printf("Names %s \n", names)

	// slice
	// expandable
	// under the hood using arrays
	// countries := make([]string, 4)

	countries := []string{}
	countries = append(countries, "India")
	countries = append(countries, "Russia")
	countries = append(countries, "China")
	countries = append(countries, "USA")
	countries = append(countries, "Africa")
	fmt.Printf("Countries %s \n", countries)
}
