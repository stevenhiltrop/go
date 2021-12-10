package main

import (
	"fmt"
	"regexp"
)

var matchNumber = regexp.MustCompile("[0-9]+")

func main() {
	fmt.Printf("%t\n", matchNumber.MatchString("Hello 123"))
}
