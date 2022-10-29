package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var toDelete string

//TODO asciiLogo prints the ascii logo of the program at the beginning
func asciLogo() {

}

//TODO outro lines with some nice messages from me
func outro() {
	fmt.Println("Thank you for using the program. Have a nice day!")
}

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
func inputGreet() {

	fmt.Println("Hello Party people! What would you like to do here?")
	_, err := fmt.Scan(&toDelete)
	if err != nil {
		return
	}
	fmt.Println("Ahh ok, keep your secrets, searching for '", toDelete, "' strings"+
		"in the database:")
}

//deleteUrl removes all entries from the history containing the given string
func deleteUrl(toDelete string) string {
	db, err := sql.Open("sqlite3", "History")

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Database closed, all entries are removed. Have a nice day!")
		}
	}(db)

	_, err = db.Exec("DELETE FROM urls WHERE url LIKE ?", "%"+toDelete+"%")

	if err != nil {
		log.Fatal(err)
	}
	return toDelete

}

// checkSqlite3Version checks what version of sqlite3 is installed on the host.
func checkV() {
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

	fmt.Println("Welcome to sqlite3 V.", version)
}

func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 1
}

func main() {

	/*fmt.Println("Let's start ...")
	future := async.Exec(func() interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running ...")
	val := future.Await()
	fmt.Println(val)*/

	checkV()
	/*deleteUrl()
	showHistory()*/

	inputGreet()

	deleteUrl(toDelete)

}
