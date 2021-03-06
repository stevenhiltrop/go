package eecummings

import (
	"regexp"
	"strings"
)

var noPunc = regexp.MustCompile("[^a-z]")

func Simplify(s string) string {
	return noPunc.ReplaceAllString(strings.ToLower(s), "")
}
