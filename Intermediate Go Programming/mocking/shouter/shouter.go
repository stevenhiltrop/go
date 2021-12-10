package shouter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

func ReadAndShout(file *os.File) error {
	scannner := bufio.NewScanner(file)

	for scannner.Scan() {
		fmt.Println(Shout(scannner.Text()))
	}

	return scannner.Err()
}
