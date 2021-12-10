package main

func main() {
	c := make(chan int)
	go func() {
		i := 0
		for {
			c <- i
			i += 1
		}
	}()
}
