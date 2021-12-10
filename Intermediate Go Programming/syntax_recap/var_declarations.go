package main

import "fmt"

func main() {
	var intSlice []int = []int{3, 1, 4, 1, 5, 9, 2, 6}
	stringSlice := []string{"The", "sky", "was", "the", "color"}

	if n := len(intSlice); n > 0 {
		fmt.Printf("intSlice has %d elemets\n", n)
	} else {
		fmt.Printf("intSlice has no elements\n")
	}

	if n := len(stringSlice); n > 0 {
		fmt.Printf("stringSlice has %d elemets\n", n)
	} else {
		fmt.Printf("stringSlice has no elements\n")
	}
}
