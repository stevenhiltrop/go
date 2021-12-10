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

type Doubler struct {
	in  <-chan int
	out chan int
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

func NewDoubler(ctx *Context, in <-chan int) *Doubler {
	d := new(Doubler)
	d.in = in
	d.out = make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := ctx.GetDone()

		for {
			select {
			case i := <-d.in:
				select {
				case d.out <- i * 2:
				case <-done:
					fmt.Printf("Doubler terminated\n")
					return
				}
			case <-done:
				fmt.Printf("Doubler terminated\n")
				return
			}
		}
	}()
	return d
}

func (c *Context) GetDone() <-chan struct{} {
	return c.done
}

func (c *Counter) GetSource() <-chan int {
	return c.c
}

func (d *Doubler) GetSource() <-chan int {
	return d.out
}

func (c *Context) Stop() {
	close(c.done)
}

func main() {
	ctx := NewContext()
	c := NewCounter(ctx)
	connect := c.GetSource()
	d := NewDoubler(ctx, connect)
	read := d.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	ctx.Stop()
	wg.Wait()
}
