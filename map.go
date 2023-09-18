package main

import "fmt"

func main() {
	// scores := make(map[string]int)

	scores := map[string]int{
		"Shihab": 3,
		"Shana":  5,
		"Shifa":  9,
	}

	// add
	scores["Althaf"] = 6

	fmt.Printf("Shihab's score: %d \n", scores["Shihab"])

	// check if exist
	shanaScore, ok := scores["Shana"]
	if !ok {
		fmt.Println("No score exist for Shana")
	} else {
		fmt.Println("Score exist for Shana and it is %d\n", shanaScore)
	}

	// delte
	delete(scores, "Shifa")
	fmt.Printf("Score %v \n", scores)

	for key, value := range scores {
		fmt.Printf("key %s: value %d\n", key, value)
	}

}
