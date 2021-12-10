package shouter

import (
	"bufio"
	"io"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

func ReadAndShout(r io.Reader) (string, error) {
	gather := ""
	scannner := bufio.NewScanner(r)

	for scannner.Scan() {
		gather += Shout(scannner.Text()) + "\n"
	}

	return gather, scannner.Err()
}
