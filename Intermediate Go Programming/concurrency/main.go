package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Context struct {
	done chan struct{}
}

type Counter struct {
	c    chan int
	done chan struct{}
	i    int
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.done = make(chan struct{})
	return ctx
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.done = make(chan struct{})

	go func() {
		for {
			select {
			case counter.c <- counter.i:
				counter.i += 1
			case <-counter.done:
				return
			}
		}
	}()

	return counter
}

func (c *Context) GetDone() <-chan struct{} {
	return c.done
}

func (c *Counter) GetSource() <-chan int {
	return c.c
}

func (c *Context) Stop() {
	close(c.done)
}

func (c *Counter) Stop() {
	c.done <- struct{}{}
}

func main() {
	c := NewCounter()
	read := c.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	c.Stop()

	fmt.Printf("%d\n", <-read)
}
