package main

import "fmt"

func main() {
	title := struct {
		English string
		French  string
	}{"Minster", "Monsieur"}
	fmt.Printf("%v\n", title)
}
