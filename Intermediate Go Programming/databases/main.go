package main

import (
	"database/sql"
	"fmt"
	"log"
	//_ "https://github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./scientists.db")
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM names")
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var first string
		var last string
		err = rows.Scan(&id, &first, &last)
		if err != nil {
			log.Fatalf("Database error:  %s", err)
		}

		fmt.Printf("%s %s\n", first, last)
	}
}
