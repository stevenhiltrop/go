package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!")

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
	}
}
