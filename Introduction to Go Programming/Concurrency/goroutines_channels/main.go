package main

func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}
	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)
}
