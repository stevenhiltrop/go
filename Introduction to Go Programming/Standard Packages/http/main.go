package main

import (
	"fmt"
	"net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["'Name"][0]
	fmt.Fprintf(w, "Poem %s coming soon\n", poemName)
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}
