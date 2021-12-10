package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error

	err = errors.New("Fatal exception")

	fmt.Printf("Error: %s\n", err)
}
