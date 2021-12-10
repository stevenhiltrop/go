package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`[
		{"English": "Mister", "French": "Monsieur"},
		{"English": "Docter", "French": "Docteur"},
		"English": "{Professer}", "French": "Professeur"}]`)

	var titles []struct {
		English string
		French  string
	}

	_ = json.Unmarshal(jsonData, &titles)

	for _, t := range titles {
		fmt.Printf("%s -> %s\n", t.English, t.French)
	}
}
