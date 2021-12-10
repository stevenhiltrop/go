package main

import (
	"eecummings"
	"flag"
	"fmt"
	"shouter"
)

var Version = "1.0.0"

func main() {
	// go run main.go -version
	version := flag.Bool("version", false, "Get version")
	flag.Parse()
	if *version {
		fmt.Printf("This is version %s\n", Version)
		return
	}
	fmt.Println(shouter.Shout("Hello, World!"))
	fmt.Println(eecummings.Simplify("Hello, World!"))
}
