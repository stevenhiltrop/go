package shouter

func TestShout(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedString := Shout(testString)
	if noLower.MatchString(shoutedString) {
		t.Fatalf("Found uppercase")
	}
}
