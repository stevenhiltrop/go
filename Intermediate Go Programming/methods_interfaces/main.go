package main

import (
	"fmt"
	"yelling"
)

func main() {
	ls := yelling.NewLoudString()
	ls.Change("Hello")
	fmt.Println(ls.String())

	n := yelling.NewLoudString
	ls2 := n()

	change := ls2.Change
	change("Be quiet")
	fmt.Println(ls2.String())

	fmt.Printf("%s\n", ls2.String())
}
