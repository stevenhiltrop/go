package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func makeBuffer() interface{} {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 200)

	var spare sync.Pool
	spare.New = makeBuffer

	for i := 0; i < 10; i++ {
		go func(offset int) {
			for {
				b := spare.Get().([]byte)
				j := offset + rand.Intn(20)
				if pool[j] != nil {
					spare.Put(pool[j])
				}
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
