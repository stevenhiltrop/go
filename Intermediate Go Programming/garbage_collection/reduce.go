package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 200)

	for i := 0; i < 10; i++ {
		go func(offset int) {
			for {
				b := makeBuffer()
				j := offset + rand.Intn(20)
				pool[j] = b
				time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
			}
		}(i * 200)
	}
	for {
		time.Sleep(1 * time.Second)
		var r runtime.MemStats
		runtime.ReadMemStats(&r)
		fmt.Printf("HeapAlloc = %d\n", r.HeapAlloc)
	}
}
