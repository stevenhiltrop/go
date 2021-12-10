package main

import (
	"fmt"
	"time"
)

type MyError struct {
	str  string
	when time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s at %s", m.str, m.when)
}

func main() {
	err := MyError{}
	err.str = "Bad thing"
	err.when = time.Now()

	fmt.Printf("%s\n", err)
}
