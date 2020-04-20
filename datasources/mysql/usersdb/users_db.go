package usersdb

import (
	"database/sql"
	"log"
	"os"

	// below blank import is needed in order for sql.Open to work since it references sqlite3.
	_ "github.com/mattn/go-sqlite3"

	// below blank import is used if connecting to a mysql server instance.  Would also need to change sql.Open to use "mysql"
	// _ "github.com/go-sql-driver/mysql"
)

const(
	mysqlUsername = "mysql_users_username"
	mysqlPassword = "mysql_users_password"
	mysqlHost 	  = "mysql_users_host"
	mysqlSchema   = "mysql_users_schema"
)

// UserDB connection for api
var (
	Client *sql.DB

	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	host 	 = os.Getenv(mysqlHost)
	schema   = os.Getenv(mysqlSchema)
)

func init() {
	// if using real connection, would use fields declared above to complete the dataSourceName and set secrets as exports
	// example: dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	dataSourceName := "./sqlite/local.db"
	var err error
	Client, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		// if there is any issue with connecting to the db, panic and do not start the app
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database succesfully configured")
}