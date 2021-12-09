package main

import (
	"fmt"
	"net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["'name"][0]
	fmt.Fprintf(w, "Poem %s coming soon\n", poemName)
}

func main() {
	// Example: curl http://127.0.0.1:8088/poem\?name=wordsworth
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}
