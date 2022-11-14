package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var toDelete string
var toOption int

//var databasePath = "" + "/History"
var storyString string

//	"C:\\Users\\stuec\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 1/History"

// TODO asciiLogo prints the ascii logo of the program at the beginning
func asciLogo() {

}

// TODO outro lines with some nice messages from me
func outro() {
	fmt.Println("Thank you for using the program. Have a nice day!")
}

// findHistory searches for the sqlite db on the pc
func findHistory() {
	/*fmt.Scan(&databasePath)


	fmt.Println(databasePath)*/

	fmt.Println("Give me your History DB file:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		databasePath := scanner.Text()
		fmt.Println("Searching for the history databasePath ...")
		fmt.Println("Input was:", databasePath)
	}

}

/*func showHistory() {
	fmt.Println(databasePath)
	db, err := sql.Open("sqlite3", databasePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT * FROM urls")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Database showed, these are your entries. Have a nice day!")
		}
	}(db)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {

		var id int
		var url string
		var title string
		var visitCount int
		var typedCount int
		var lastVisitTime int
		var hidden int

		err = rows.Scan(&id, &url, &title, &visitCount, &typedCount, &lastVisitTime, &hidden)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, url, title, visitCount, typedCount, lastVisitTime, hidden)

		// and then print out the id, url, title, visit_count, typed_count, last_visit_time, hidden only with fmt.Println
		fmt.Println(id, url, title, visitCount, typedCount, lastVisitTime, hidden)
	}
}*/

func checkByDelete() {

	/*db, err := sql.Open("sqlite3", databasePath)

	if err != nil {
		log.Fatal(err)

	}

	rows, err := db.Query("SELECT * FROM urls WHERE url LIKE ?", "%"+toDelete+"%")

	if err != nil {

		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Database showed, these are your entries. Have a nice day!")
		}
	}(db)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {

		var id int
		var url string
		var title string
		var visitCount int
		var typedCount int
		var lastVisitTime int
		var hidden int

		err = rows.Scan(&id, &url, &title, &visitCount, &typedCount, &lastVisitTime, &hidden)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("You sure you want to delete?")

		// and then print out the id, url, title, visit_count, typed_count, last_visit_time, hidden only with fmt.Println
		fmt.Println(id, url, title, visitCount, typedCount, lastVisitTime, hidden)
	}*/
}

// TODO
// greetDelete window asks for the input and option
func greetDelete() {

	fmt.Println("Hello Party people! What keyword would you like to delete From your history?")
	_, err := fmt.Scan(&toDelete)
	if err != nil {
		return
	}

	fmt.Println("Ahh ok, keep your secrets, searching for '", toDelete, "' strings"+
		" in the databasePath:")

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

	checkByDelete()
	fmt.Println("Are you sure you want to delete these entries? (1 = yes, 2 = no)")
	fmt.Scan(&toOption)
	if toOption == 1 {
		_, err = db.Exec("DELETE FROM urls WHERE url LIKE ?", "%"+toDelete+"%")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Ok, no entries deleted. Have a nice day!")
	}

	return toDelete

}

// checkV checks what version of sqlite3 is installed on the host.
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

func actionSelect() int {
	fmt.Println("Let's start ... \nWhat would you like to do?\n 1. Show history\n 2. Delete history")
	fmt.Scan(&toOption)
	// loop to check for the right input
	for toOption != 1 && toOption != 2 {
		fmt.Println("Please enter a valid option:\n 1. Show history\n 2. Delete history")
		fmt.Scan(&toOption)
	}
	return toOption
}

// control flow
func controlFlow() {
	// toOption switch case for the options to show and delete history
	switch toOption := toOption; {
	case toOption == 1:
		fmt.Println("one")
		//showHistory()
	case toOption == 2:
		fmt.Println("two")
		//greetDelete()
		//deleteUrl(toDelete)
	default:
		fmt.Println("Invalid option, try again")
	}
}

func main() {
	checkV()
	findHistory()
	actionSelect()
	//controlFlow()

}

func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 1
}

// async function
/*fmt.Println("Let's start ...")
future := async.Exec(func() interface{} {
	return DoneAsync()
})
fmt.Println("Done is running ...")
val := future.Await()
fmt.Println(val)*/
