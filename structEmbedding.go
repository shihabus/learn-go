package main

import "fmt"

// struct embedding
// it is kind of inheritance
type Position struct {
	x, y int
}

func (p *Position) Move(val int) {
	fmt.Println("The position is moved by: ", val)
}

type Entity struct {
	Position
	id   string
	name string
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := SpecialEntity{
		specialField: 10.01,
		Entity: Entity{
			id:   "001",
			name: "Pokemon",
			Position: Position{
				x: 13,
				y: 34,
			},
		},
	}
	e.x = 30 //
	e.Move(200)
	fmt.Printf("Entity is: %+v%", e)
}
