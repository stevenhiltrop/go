package main

import "fmt"

func main() {
	nums := []int{1, 11, 21, 1211, 111221, 312211}

	middle := nums[1:3]
	fmt.Printf("%v\n", middle)

	nums[1] *= 2
	fmt.Printf("%v\n", middle)

	middle = append(middle, 42)
	fmt.Printf("%v\n", middle)
	fmt.Printf("%v\n", nums)
}
