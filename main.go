package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func showHistory() {
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

		fmt.Println(id, url, title, visit_count, typed_count, last_visit_time, hidden)

		// and then print out the id, url, title, visit_count, typed_count, last_visit_time, hidden only with fmt.Println
		fmt.Println(id, url, title, visit_count, typed_count, last_visit_time, hidden)
	}
}

//TODO
//hello window asks for the input and option
func hi() {
	var toDelete string

	fmt.Println("Hello Party people! What would you like to delete here jiggo?")
	scan, err := fmt.Scan(&toDelete)
	if err != nil {
		return
	}
	fmt.Println("It's ok homie, keep your secrets, all rows with the keyword", scan, "will be deleted")
}

//deleteUrl removes all entries from the history containing the given string
func deleteUrl(toDelete string) string {
	db, err := sql.Open("sqlite3", "History")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM urls WHERE url LIKE ?", "%"+toDelete+"%")

	if err != nil {
		log.Fatal(err)
	}
	return toDelete

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
	//checkSqlite3Version()
	//deleteUrl()
	//showHistory()
	var toDelete string

	fmt.Println("Hello Party people! What would you like to delete here jiggo?")
	fmt.Scan(&toDelete)

	fmt.Println("You want to delete", toDelete)
	deleteUrl(toDelete)
}
