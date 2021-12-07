package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

type intSlice []int
type stringSlice []string

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

func (i intSlice) Len() int {
	return len(i)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (i stringSlice) Len() int {
	return len(i)
}

func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6}
	shuffle(is)
	fmt.Printf("%v\n", is)

	ss := stringSlice{"The", "quick", "brown", "fox"}
	shuffle(ss)
	fmt.Printf("%v\n", ss)
}
