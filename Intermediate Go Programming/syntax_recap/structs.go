package main

import "fmt"

type Translation struct {
	English string
	French  string
}

func main() {
	title := Translation{"Minster", "Monsieur"}
	fmt.Printf("%v\n", title)
}
