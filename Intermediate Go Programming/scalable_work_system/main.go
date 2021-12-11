package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var n int

type task interface {
	process()
	output()
}

type factory interface {
	create(line string) task
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			in <- f.create(s.Text())
		}

		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.output()
	}
}

type HTTPTask struct {
	url string
	ok  bool
}

func (h *HTTPTask) process() {
	if h.url == "" {
		h.ok = false
		return
	}
	resp, err := http.Get(h.url)
	if err != nil {
		h.ok = false
		return
	}

	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false
}

func (h *HTTPTask) output() {
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type Factory struct{}

func (f *Factory) create(line string) task {
	h := &HTTPTask{}
	h.url = line
	return h
}

func main() {
	// echo "https://www.cloudflare.com | ./scalable_work_system"
	// ./scalable_work_system < urls.txt
	count := flag.Int("count", 10, "Number of workers")
	flag.Parse()
	n = *count
	f := &Factory{}
	run(f)
}
