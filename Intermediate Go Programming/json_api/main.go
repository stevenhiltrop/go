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

func marshal(w http.ResponseWriter, i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(j)
}

func (ds *databaseServer) get(w http.ResponseWriter, req *http.Request, last string) {
	rows, err := ds.db.Query("SELECT * FROM names WHERE last = ?;", last)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var result struct {
		Success bool
		Id      int
		First   string
		Last    string
		Alive   bool
	}
	if rows.Next() {
		result.Success = true
		rows.Scan(&result.Id, &result.First, &result.Last, &result.Alive)
		marshal(w, result)
	} else {
		var result struct {
			Succes bool
		}
		marshal(w, result)
	}
}

func (ds *databaseServer) post(w http.ResponseWriter, req *http.Request, last string) {
	if req.Form["first"] == nil {
		http.Error(w, fmt.Sprintf("Missing first"), http.StatusBadRequest)
		return
	}

	first := req.Form["first"][0]

	_, err := ds.db.Exec("INSERT INTO names (first, last, alive) VALUES (?, ?, 1)", first, last)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		return
	}

	var result struct {
		Success bool
	}
	result.Success = true
	marshal(w, result)
}

func (ds *databaseServer) api(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Form["last"] == nil {
		http.Error(w, fmt.Sprintf("Missing last"), http.StatusBadRequest)
		return
	}

	last := req.Form["last"][0]

	switch {
	case req.Method == "GET":
		ds.get(w, req, last)
	case req.Method == "POST":
		ds.post(w, req, last)
	default:
		http.Error(w, "Bad method", http.StatusMethodNotAllowed)
	}
}

func main() {
	/* CLI:

	GET: curl -s http://127.0.0.1:8080/api\?last=Lovelace | json_pp
	POST: curl -s -X POST http://127.0.0.1:8080/api\?last=Sommerville\&first=Mary | json_pp
	*/
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
