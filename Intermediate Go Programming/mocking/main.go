package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"shouter"
)

func main() {
	filename := flag.String("file", "", "File to shout")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Failed to open %s", *filename)
		return
	}
	s, _ := shouter.ReadAndShout(file)
	fmt.Println(s)
}
