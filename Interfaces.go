package main

import (
	"fmt"
)

// interface
// basically a contract
// internal implementation is abstracted
type NumStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

// implementation 1
type MongoDBStorer struct{}

func (m MongoDBStorer) GetAll() ([]int, error) {
	return []int{1, 2, 3, 5}, nil
}

func (m MongoDBStorer) Put(number int) error {
	fmt.Println("Successfully stored ", number)
	return nil
}

// implementation 2
type PostgresStorer struct{}

func (p PostgresStorer) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8}, nil
}

func (p PostgresStorer) Put(number int) error {
	fmt.Println("Successfully stored ", number)
	return nil
}

type ApiServer struct {
	numberStore NumStorer
}

func main() {
	// as you see with interface it is
	// easier to switch over the implementation
	// as far as the ApiServer.numberStore is concerned
	// it just need two internal methods
	// that satisfy the func signature in the interface
	apiServer := ApiServer{
		// numberStore: MongoDBStorer{},
		numberStore: PostgresStorer{},
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
