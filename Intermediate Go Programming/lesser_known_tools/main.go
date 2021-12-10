package main

import "fmt"

func Printf(s string) {
	fmt.Println(s)
}

func main() {
	/* CLI:
	go vet main.go
	/bin/golint lesser_known_tools
	go list main.go | more
		go list -f "{{ .Imports }} lesser_known_tools"
		go list -f "{{ .Deps }} lesser_known_tools"
		go list -f "{{ .Dir }} lesser_known_tools"
		go list -f "{{ .GoFiles }} lesser_known_tools"
	*/
	Printf("Hello, World! %s")

	return
}
