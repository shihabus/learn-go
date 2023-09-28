package main

import (
	"fmt"
	"strings"
)

// typed functions
type TransformFunc func(s string) string

func TransformStr(s string, fn TransformFunc) string {
	return fn(s)
}

func UppercaseStr(s string) string {
	return strings.ToUpper(s)
}

// function decorator 
func Prefixr(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + "_" + s
	}
}

func main() {
	fmt.Println(TransformStr("Upper cased string", UppercaseStr))
	fmt.Println(TransformStr("World", Prefixr("Hello")))
	fmt.Println(TransformStr("Foo", Prefixr("Bar")))
}
