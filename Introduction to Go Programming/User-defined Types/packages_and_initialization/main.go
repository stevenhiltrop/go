package main

import (
	"fmt"
	"shuffler"
)

type stringSlice []weightString
type weightString struct {
	weight int
	s      string
}

func (is stringSlice) Len() int {
	return len(is)
}

func (is stringSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	i := stringSlice{weightString{100, "Hello"}, weightString{200, "World!"}, weightString{10, "Goodbye"}}
	shuffler.Shuffle(i)
	fmt.Printf("%v", i)
	fmt.Printf("Loop: %d", shuffler.GetCount())
}
