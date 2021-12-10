package main

import (
	"fmt"
)

func main() {
	var (
		even   int
		odd    int
		zeroes int
		total  int
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, -3, 10}

IgnoreNegative:
	for _, n := range numbers {
		switch {
		case n == 0:
			zeroes += 1
		case n%2 == 0:
			even += 1
		case n%2 == 1:
			odd += 1
		default:
			continue IgnoreNegative
		}
		total += 1
	}

	fmt.Printf("Even %d, odd %d, zeroes %d, total %d/%d\n", even, odd, zeroes, total, len(numbers))
}
