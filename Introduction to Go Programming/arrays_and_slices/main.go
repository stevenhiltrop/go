package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := [4]string{"the", "quick", "brown", "fox"}

	printer(words[:])
}
