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
	ctx *Context
	c   chan int
	i   int
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.done = make(chan struct{})
	return ctx
}

func NewCounter(ctx *Context) *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := counter.ctx.GetDone()

		for {
			select {
			case counter.c <- counter.i:
				counter.i += 1
			case <-done:
				fmt.Printf("Counter terminated\n")
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

func main() {
	ctx := NewContext()
	c := NewCounter(ctx)
	read := c.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	ctx.Stop()
	wg.Wait()
}
