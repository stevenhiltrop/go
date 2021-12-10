package main

import (
	"flag"
	"fmt"
	"time"
)

/* User defined time problem:
var timeout = 1000 // Milliseconds

Solution: use the time module
var timeout = 1 * time.Millisecond
*/
func main() {
	timeout := flag.Duration("timeout", 1*time.Microsecond, "Timeout value")
	flag.Parse()

	fmt.Printf("Timeout us %s\n", *timeout)
}
