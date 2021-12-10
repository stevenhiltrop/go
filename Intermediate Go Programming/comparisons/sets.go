package main

import "fmt"

type Person struct {
	First, Last string
}

func main() {
	people := make(map[Person]bool)

	people[Person{"John", "Graham-Cumming"}] = true
	people[Person{"Justin", "Bieber"}] = false

	t, ok := people[Person{"Adam", "Smith"}]
	if ok {
		fmt.Printf("%t\n", t)
	} else {
		fmt.Printf("Not present")
	}
}
