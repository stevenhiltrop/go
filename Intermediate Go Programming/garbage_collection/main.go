package main

import (
	"fmt"
	"runtime"
	"time"
)

func length(s string) int {
	c := []byte(s)
	n := len(c)
	b := make([]byte, n)
	return len(b)
}

func main() {
	s := "Hello, World!"
	s += s
	s += s
	s += s

	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap size %d\n", r.HeapAlloc)
			fmt.Printf("NumGC %d\n", r.NumGC)
			start = time.Now()
		}
		length(s)
	}
}
