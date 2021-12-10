package main

import (
	"fmt"
)

func do() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("An error occured: %s\n", err)
		}
	}()
	err = fmt.Errorf("Bad thing happened")
}

func main() {
	do()
}
