package main

import (
	"database/sql"
	"fmt"
	"log"
	//_ "https://github.com/mattn/go-sqlite3"
)

func show(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM names")
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var first string
		var last string
		var alive bool
		err = rows.Scan(&id, &first, &last, &alive)
		if err != nil {
			log.Fatalf("Database error:  %s", err)
		}

		fmt.Printf("%s %s %t\n", first, last, alive)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./scientists.db")
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}
	defer db.Close()

	show(db)

	t, err := db.Begin()
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}

	st, err := t.Prepare("UPDATE names SET alive = 0;")
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}

	_, err = st.Exec()
	if err != nil {
		log.Fatalf("Database error:  %s", err)
	}

	show(db)

	t.Commit()

	show(db)
}
