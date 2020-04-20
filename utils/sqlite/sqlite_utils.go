package sqlite

import (
    "database/sql"
    "fmt"
    "strconv"
    _ "github.com/mattn/go-sqlite3"
)
// CreateDb ...
func CreateDb() {
    database, _ := sql.Open("sqlite3", "./database/local.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, email TEXT, datecreated TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO users (id, firstname, lastname, email, datecreated) VALUES (?, ?, ?, ?, ?)")
    statement.Exec(1, "sdet", "automation", "sdet.testautomation@gmail.com", "2020-04-19-T22:58:14Z")
    rows, _ := database.Query("SELECT id, firstname, lastname, email, datecreated FROM users")
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