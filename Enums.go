package main

import (
	"fmt"
)

// Enum
type WeaponType int

const (
	Axe WeaponType = iota // will increment the following each by 1
	Sword
	Knife
	Gun
)

// Override the value to be shown when being printed
func (w WeaponType) String() string {
	var weaponName string
	switch w {
	case Axe:
		weaponName = "Axe"
	case Sword:
		weaponName = "Sword"
	case Knife:
		weaponName = "Knife"
	case Gun:
		weaponName = "Gun"
	}
	return weaponName
}

func (w WeaponType) getDamage() int {
	switch w {
	case Axe:
		return 60
	case Sword:
		return 40
	case Knife:
		return 20
	case Gun:
		return 80
	default:
		panic("No weapon found")
	}
}

func main() {
	fmt.Println("Gun", Gun)
	fmt.Println("Damage of Knife", Knife.getDamage())
}
