package main

import "fmt"

// POINTER
// 8 bytes long
// carries the memory address
// when parameter are passed,
// they are passed by value (copies are being passed)
// with pointers we can pass them by reference
// used for
// - state modification/update
// - when we have to pass large data. While being passed by value, the data have to be copied/duplicated
type Player struct {
	health int
}

func takeDamage(p Player, damage int) {
	p.health -= damage
	fmt.Printf("Player was attacked, and current health is %v\n", p.health)
}

// with pointer
func takeDamagePinter(p *Player, damage int) {
	p.health -= damage
	fmt.Printf("Player was attacked, and current health is %v\n", p.health)
}

func main() {
	player1 := Player{
		health: 100,
	}
	// the changed value is only scoped with the function
	// outide the function health is still 100
	takeDamage(player1, 3)  // 97
	takeDamage(player1, 10) // 90

	fmt.Printf("player1 %+v\n", player1) // 100

	// pass by regference to fix this
	player2 := &Player{
		health: 100,
	}
	takeDamagePinter(player2, 3)  // 97
	takeDamagePinter(player2, 10) // 87

	fmt.Printf("Player2 %+v\n", player2) // 87
}
