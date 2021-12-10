package main

import "fmt"

func main() {
	var (
		even  int
		odd   int
		total int
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, 10}

Abort:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			total += 1

			if n == 0 {
				break Abort
			}

			if n%2 == 0 {
				even += 1
			} else {
				odd += 1
			}
		}
	}

	fmt.Printf("Even %d, odd %d, total %d", even, odd, total)
}
