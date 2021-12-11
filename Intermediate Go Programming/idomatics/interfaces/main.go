package main

import (
	"fmt"
	"io"
	"os"
)

type Pair struct {
	x, y string
}

func (p Pair) String() string {
	return fmt.Sprintf("x=%s, y=%s", p.x, p.y)
}

// Reading file problem: Don't use *os.File when handling files
func (p Pair) LoadY(f *os.File) error {
	return nil
}

// Solution: use io.Reader
func (p Pair) LoadX(r io.Reader) error {
	return nil
}

func main() {
	p := Pair{x: "XXX", y: "YYY"}
	fmt.Printf("%s\n", p)
}
