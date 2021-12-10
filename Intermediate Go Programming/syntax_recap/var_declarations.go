package main

import "fmt"

func main() {
	var intSlice []int = []int{3, 1, 4, 1, 5, 9, 2, 6}
	stringSlice := []string{"The", "sky", "was", "the", "color"}

	_, err := fmt.Printf("intSlice has %d elements\n", len(intSlice))
	_, err = fmt.Printf("stringSlice has %d elements\n", len(stringSlice))

	fmt.Printf("%s\n", err)
}
