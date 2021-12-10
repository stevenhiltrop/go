package main

import (
	"fmt"
	"yelling"
)

func main() {
	ls := yelling.NewLoudString()
	ls.Change("Hello")
	fmt.Println(ls.String())
}
