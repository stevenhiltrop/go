package shouter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

func ReadAndShout(r io.Reader) error {
	scannner := bufio.NewScanner(r)

	for scannner.Scan() {
		fmt.Println(Shout(scannner.Text()))
	}

	return scannner.Err()
}
