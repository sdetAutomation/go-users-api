package sqlitehelper

import (
	"database/sql"
	"fmt"
	"strconv"

	// below blank import is needed in order for CreateDb func to work since line 12 references sqlite3.
	_ "github.com/mattn/go-sqlite3"
)

// CreateDb ...
func CreateDb() {
	database, _ := sql.Open("sqlite3", "./sqlite/local.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY NOT NULL, first_name TEXT, last_name TEXT, email TEXT NOT NULL UNIQUE, date_created TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)")
	statement.Exec("sdet", "automation", "sdet.testautomation@gmail.com", "2020-04-19-T22:58:14Z")
	rows, _ := database.Query("SELECT id, first_name, last_name, email, date_created FROM users")
	var id int
	var firstname string
	var lastname string
	var email string
	var datecreated string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &email, &datecreated)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + email + " " + datecreated)
	}
}
