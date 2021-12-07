package main

import "fmt"

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}

	i := 0

	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done:
			close(done)
			return
		}
	}
}

func main() {
	wordChan := make(chan string)
	doneChan := make(chan bool)

	go emit(wordChan, doneChan)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordChan)
	}

	doneChan <- true
}
