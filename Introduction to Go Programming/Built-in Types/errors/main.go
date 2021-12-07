package main

import (
	"fmt"
	"os"
)

func printer(msg string) error {
	if msg == "" {
		return fmt.Errorf("No message was parsed")
	}

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {
	if err := printer(""); err != nil {
		os.Exit(1)
	}
}
