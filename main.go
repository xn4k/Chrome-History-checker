package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func execRow() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
		SELECT 
		  datetime(last_visit_time/1000000-11644473600, "unixepoch") as last_visited, 
		  url, 
		  title, 
		  visit_count 
		FROM urls
		`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table cars created")
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

}
