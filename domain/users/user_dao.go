package users

// data-access-object - access layer to the database. Persist and retrieve user from database.  
// this file is the only place in the app that can access the db.
// this pattern allows for easy management and switching between persistant data bases.   

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sdetAutomation/go-users-api/datasources/mysql/usersdb"
	"github.com/sdetAutomation/go-users-api/utils/date"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get ...
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DataCreated); getErr != nil {
		if strings.Contains(getErr.Error(), errorNoRows) {
			return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error retrieving user %d record: %s", user.ID, getErr.Error()))
	}
	return nil
}

// Save ...
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// defer will execute right before a return statement is executed.
	defer stmt.Close()

	user.DataCreated = date.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DataCreated)
	if saveErr != nil {
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", saveErr.Error()))
		}
		// // below can tell you what error numbers and messages. 
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Number)
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", saveErr.Error()))	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", err.Error()))
	}
	// update user with last insert id. 
	user.ID = userID
	return nil
}