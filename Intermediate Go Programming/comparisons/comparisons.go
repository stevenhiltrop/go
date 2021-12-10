package main

import "fmt"

type MyString struct {
	Str string
}

type MyInterface interface {
	String() string
}

type Person struct {
	First, Last string
}

func main() {
	// Integers
	var i int = 42
	var j int = 43 // int8

	fmt.Printf("%t\n", i == j)

	// Strings
	s := "Hello"
	t := "World!"

	fmt.Printf("%t\n", s == t)

	// Structs
	x := MyString{"Hello"}
	y := MyString{"World!"}

	fmt.Printf("%t\n", x == y)

	// Interfaces
	var a MyInterface
	var b MyInterface

	fmt.Printf("%t\n", a == b)

	// Maps
	people := make(map[Person]bool)

	people[Person{"John", "Graham-Cumming"}] = true
	people[Person{"Justin", "Bieber"}] = false

	fmt.Printf("%#v\n", people)
}
