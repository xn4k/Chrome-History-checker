package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func execRow() {
	db, err := sql.Open("sqlite3", "History")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM urls")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var url string
		var title string
		var visit_count int
		var typed_count int
		var last_visit_time int
		var hidden int

		err = rows.Scan(&id, &url, &title, &visit_count, &typed_count, &last_visit_time, &hidden)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s %d\n", id, url, title, visit_count, typed_count, last_visit_time, hidden)
	}
}

// checkSqlite3Version checks what version of sqlite3 is installed on the host.
func checkSqlite3Version() {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func main() {
	execRow()

}
