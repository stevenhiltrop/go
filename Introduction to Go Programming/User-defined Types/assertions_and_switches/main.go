package main

import "fmt"

func whatIsThis(i interface{}) {
	//fmt.Printf("%T\n", i)
	switch i.(type) {
	case string:
		s := i.(string)
		fmt.Printf("It's a string %s\n", s) // Better
	case uint32:
		fmt.Printf("It's an unsigned 32-bit integer %d\n", i.(uint32)) // Bad
	default:
		fmt.Printf("Don't know\n")
	}
}

func main() {
	whatIsThis([...]string{"A", "B", "C"})
}
