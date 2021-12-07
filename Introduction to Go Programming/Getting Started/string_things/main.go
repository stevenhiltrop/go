package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	for _, r := range atoz {
		fmt.Printf("%c\n", r)
	}
}
