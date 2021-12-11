package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type databaseServer struct {
	db *sql.DB
}

func (ds *databaseServer) api(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Form["last"] == nil {
		http.Error(w, fmt.Sprintf("Missing last"), http.StatusBadRequest)
		return
	}

	last := req.Form["last"][0]

	rows, err := ds.db.Query("SELECT * FROM names WHERE last = ?;", last)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	if rows.Next() {
		var result struct {
			Id    int
			First string
			Last  string
			Alive bool
		}
		rows.Scan(&result.Id, &result.First, &result.Last, &result.Alive)

		j, err := json.Marshal(&result)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
			return
		}
		w.Write(j)
	} else {
		fmt.Fprintf(w, "%s not found\n", last)
	}
}

func main() {
	ds := &databaseServer{}
	var err error
	ds.db, err = sql.Open("sqlite3", "./scientists.db")
	if err != nil {
		log.Fatalf("Database error: %s", err)
	}
	defer ds.db.Close()

	http.HandleFunc("/api", ds.api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
