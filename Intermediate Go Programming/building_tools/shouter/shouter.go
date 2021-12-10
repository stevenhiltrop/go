package shouter

import "strings"

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}
