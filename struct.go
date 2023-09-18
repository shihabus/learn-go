package main

import "fmt"

type Player struct {
	name   string
	health int
	power  string
}

func (p Player) getHealth() int {
	return p.health
}

func main() {
	shihab := Player{
		name:   "Shihab",
		health: 80,
		power:  "Kick",
	}

	fmt.Printf("%s has %d health", shihab.name, shihab.getHealth())
}
