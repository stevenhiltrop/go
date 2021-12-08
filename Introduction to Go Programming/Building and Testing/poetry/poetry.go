package poetry

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func (p Poem) Stats() (numVowels, numConsonants int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels += 1
				default:
					numConsonants += 1
				}
			}
		}
	}
	return
}
