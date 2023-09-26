package main

import "fmt"

type Color int

// a way to return proper strings
// while printing the values
// fmt will call this method
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "Blue"
	case ColorBlack:
		return "Black"
	case ColorYellow:
		return "Yellow"
	case ColorPink:
		return "Pink"

	default:
		panic("Invalid color")
	}
}

// a way to define enums
const (
  // iota will increment 
  // everything below it
  // usually starts at 0
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Println("The color I love is", ColorBlack)
}
