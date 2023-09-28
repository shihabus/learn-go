package main

import "fmt"

// a custom map that has
// string key and int values
type CustomMap struct {
	data map[string]int
}

// a method to insert value into the map
func (m *CustomMap) Insert(k string, val int) error {
	m.data[k] = val
	return nil
}

// used to initialize a customMap
func NewCustomMap() *CustomMap {
	return &CustomMap{
		data: make(map[string]int),
	}
}

// Problem:
// here there is tight binding
// the custom map can onlys store
// int values with string keys
// if we want to store string values
// we are forced to create a new map

// ---- GENERICS -------
type GenericCustomMap[K comparable, V any] struct {
	data map[K]V
}

func (g *GenericCustomMap[K, V]) Insert(k K, val V) error {
	g.data[k] = val
	return nil
}

func NewGenericMap[K comparable, V any]() *GenericCustomMap[K, V] {
	return &GenericCustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	userAge := NewCustomMap()
	userAge.Insert("Shihab", 30)
	userAge.Insert("Shana", 25)
	userAge.Insert("Haala", 1)
	fmt.Printf("User age %+v\n", userAge)

	userId := NewGenericMap[string, int]()
	userId.Insert("Shihab", 1)
	userId.Insert("Shana", 2)
	userId.Insert("Haala", 3)
	fmt.Printf("User id %+v\n", userId)

	userName := NewGenericMap[string, string]()
	userName.Insert("Shihab", "sus")
	userName.Insert("Shana", "shanu")
	userName.Insert("Haala", "azlin")
	fmt.Printf("User names %+v\n", userName)
}
