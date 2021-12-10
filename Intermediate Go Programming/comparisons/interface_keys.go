package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var output = make(map[io.Writer]bool)

func main() {
	output[os.Stdout] = true
	output[os.Stderr] = false

	b := new(bytes.Buffer)
	output[b] = true

	for w, enabled := range output {
		if enabled {
			fmt.Fprintf(w, "Some output \n")
		}
	}
	fmt.Printf("%d bytes in buffer\n", b.Len())
}
