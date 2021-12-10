package main

import "fmt"

func main() {
	var x interface{}

	if _, ok := x.(interface {
		String() string
	}); !ok {
		fmt.Printf("x can't String()")
	}
}
