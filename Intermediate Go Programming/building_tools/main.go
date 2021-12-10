package main

import (
	"eecummings"
	"fmt"
	"shouter"
)

func main() {
	fmt.Println(shouter.Shout("Hello, World!"))
	fmt.Println(eecummings.Simplify("Hello, World!"))
}
