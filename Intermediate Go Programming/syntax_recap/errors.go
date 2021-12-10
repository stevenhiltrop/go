package main

import (
	"errors"
	"fmt"
)

var (
	ErrBadStartup = errors.New("Failed to start correctly")
)

func main() {
	var err error
	err = ErrBadStartup

	if err == ErrBadStartup {
		//IGNORE
		return
	}
	fmt.Printf("Error: %s\n", err)
}
