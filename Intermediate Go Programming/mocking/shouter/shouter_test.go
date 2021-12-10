package shouter

import (
	"regexp"
	"strings"
	"testing"
)

func TestShout(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedString := Shout(testString)
	if noLower.MatchString(shoutedString) {
		t.Fatalf("Found lowercase")
	}
}

func TestReadAndShout(t *testing.T) {
	testString := "Hello\nWorld\n"

	s, err := ReadAndShout(strings.NewReader(testString))

	if err != nil {
		t.Fatalf("Error ReadAndShout: %s", err)
	}

	noLower := regexp.MustCompile("[a-z]")
	if noLower.MatchString(s) {
		t.Fatalf("Found lowercase")
	}

}
