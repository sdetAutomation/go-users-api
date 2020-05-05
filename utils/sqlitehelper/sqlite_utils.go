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
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY NOT NULL, first_name TEXT, last_name TEXT, email TEXT NOT NULL UNIQUE, date_created DATETIME, status TEXT NOT NULL, password TEXT NOT NULL)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?)")
	statement.Exec("sdet", "automation", "sdet.testautomation@gmail.com", "2020-04-19 22:58:14", "active", "il0veg0")
	rows, _ := database.Query("SELECT id, first_name, last_name, email, date_created, status FROM users")
	var id int
	var firstname string
	var lastname string
	var email string
	var datecreated string
	var status string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &email, &datecreated, &status)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + email + " " + datecreated + " " + status)
	}
}
