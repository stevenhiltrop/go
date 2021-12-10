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
		err    error
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, -3, 10}

Abort:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			total += 1

			switch {
			case n == 0:
				zeroes += 1
			case n%2 == 0:
				even += 1
			case n%2 == 1:
				odd += 1
			default:
				break Abort
			}
			total += 1
		}
	}

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	fmt.Printf("Even %d, odd %d, zeroes %d, total %d", even, odd, zeroes, total)
}
