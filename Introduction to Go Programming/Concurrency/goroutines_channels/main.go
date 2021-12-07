package main

import (
	"fmt"
	"math/rand"
)

func emitString(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}
	close(c)
}

func emitRandoms(c chan int) {

	for {
		c <- rand.Intn(1000)
	}
}

func makeID(idChan chan int) {
	var id int
	id = 0

	for {
		idChan <- id
		id += 1
	}
}

func main() {
	// Strings
	/* wordChannel := make(chan string)

	go emitString(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	} */

	// Integers
	/* randoms := make(chan int)

	go emitRandoms(randoms)

	for n := range randoms {
		fmt.Printf("%d ", n)
	} */

	idChan := make(chan int)

	go makeID(idChan)

	fmt.Printf("ID #%d\n", <-idChan)
	fmt.Printf("ID #%d\n", <-idChan)
	fmt.Printf("ID #%d\n", <-idChan)
	fmt.Printf("ID #%d\n", <-idChan)
}
