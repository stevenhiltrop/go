package main

import (
	"fmt"
)

func main() {
	p := []string{
		"And from my pillow, looking forth by light",
		"Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood",
		"Of Newton with his prism and silent face",
		"The marble index of a mind for ever",
		"Voyagin through strange seas of Thought, alone."}

	fmt.Printf("Type: %T\n", p[0])

	result := ""

	for _, s := range p {
		result += fmt.Sprintf("%s\n", s)
	}

	fmt.Printf(result)
}
