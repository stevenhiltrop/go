package main

import "fmt"

func printer(msg, msg2 string) {
	fmt.Printf("%s\n", msg)
	fmt.Printf("%s\n", msg2)
}

func main() {
	printer("Hello,", " World!")
}
