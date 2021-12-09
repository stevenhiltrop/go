package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line string
type Stanza []Line
type Poem []Stanza

func loadPoem(name string) (Poem, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Poem{}

	var s Stanza

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		l := scan.Text()
		if l == "" {
			p = append(p, s)
			s = Stanza{}
			continue
		}
		s = append(s, Line(l))
	}
	p = append(s, s)

	if scan.Err() != nil {
		return nil, scan.Err()
	}
}

func main() {
	p, err := loadPoem("wordsworth")
	if err != nil {
		fmt.Printf("Error loading poem: %s\n", err)
	}
	fmt.Printf("%s\n", p)

	/*
		p := []string{
			"And from my pillow, looking forth by light",
			"Of moon or favouring stars, I could behold",
			"The antechapel where the statue stood",
			"Of Newton with his prism and silent face",
			"The marble index of a mind for ever",
			"Voyagin through strange seas of Thought, alone."}
	*/
}
