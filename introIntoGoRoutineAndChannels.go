package main

import (
	"fmt"
	"time"
)

func main() {
	// will take time to run
	// result := fetchSomeResource()
	// fmt.Println("The result: ", result)

	// go routine
	// these are tasks scheduled by go runtime
	// to be executed later

	// make it async
	go fetchSomeResource()
	fmt.Println("We are not blocked")

	// CHANNELS
	// channels are like pipe
	// you input <- in one end
	// and recieve <- in the other end
	// there are basically two types of channel

	// 1 unbuffered
	// resultch:=make(chan string)

	// 2 buffered channel
	// here 10 is the size of channel
	// so it can store 10 strings
	// resultch:=make(chan string,10)

	// one thing to understand is
	// code execution gets blocked when a channel is full
	// so the code below it is never executed

	// approach 1 (unbuffered channel)
	resultch1 := make(chan string)

	// this work because,
	// the code block is being scheduled
	// to be executed later
	go func() {
		result1 := <-resultch1
		fmt.Println("The result1 is: ", result1)
	}()

	resultch1 <- "foo"

	// approach 2 (using buffered channel)
	resultch2 := make(chan string, 1)
	resultch2 <- "bar"
	result2 := <-resultch2
	fmt.Println("The result2 is: ", result2)

}

func fetchSomeResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}
