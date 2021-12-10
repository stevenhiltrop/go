package main

import (
	"bufio"
	"fmt"
	"lru1"
	"os"
	"runtime/pprof"
)

func main() {
	/* CLI:
	go build lrutest.go
	cat text_file | ./lrutest
	cat wordlist | ./lrutest
	go tool pprof lrutest lrutest.cpuprofile
	*/
	f, _ := os.Create("lrutest.cpuprofile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	cache := lru1.NewCache(1000)

	count := 0
	miss := 0

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		word := in.Text()
		if len(word) > 3 {
			three := word[1:4]
			if cache.Get(three) == nil {
				cache.Put(three, word)
				miss += 1
			}
			count += 1
		}
	}
	fmt.Printf("%d total %d misses\n", count, miss)
}
