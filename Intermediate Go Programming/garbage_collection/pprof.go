package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func save(sizes []int, n int) []int {
	return append(sizes, n)
}

func main() {
	// CLI: go tool pprof pprof profile
	// (pprof) web
	// (pprof) top
	s := "Hello, World!"
	var sizes []int
	first := true
	start := time.Now()

	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap size %d\n", r.HeapAlloc)
			fmt.Printf("NumGC %d\n", r.NumGC)
			start = time.Now()
			if first {
				out, _ := os.Create("profile")
				pprof.WriteHeapProfile(out)
				out.Close()
				first = false
			}
		}
		sizes = save(sizes, len(s))
	}
}
