package main

import "fmt"

func empty(i []int) (n int) {
	if n := len(i); n > 0 {
		n = 1
	}
	return
}

func main() {
	var intSlice []int = []int{3, 1, 4, 1, 5, 9, 2, 6}

	fmt.Printf("%d\n", empty(intSlice))
}
