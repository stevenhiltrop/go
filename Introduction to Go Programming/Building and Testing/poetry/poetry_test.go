package poetry

import "testing"

func TestNumStanzas(t *testing.T) {
	p := Poem{}
	if p.NumStanzas() != 0 {
		t.Fatalf("Empty poem is not empty!")
	}

	p = Poem{{
		"And from my pillow, looking forth by light",
		"Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood",
		"Of Newton with his prism and silent face",
		"The marble index of a mind for ever",
		"Voyagin through strange seas of Thought, alone."}}

	if p.NumStanzas() != 6 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}
}
