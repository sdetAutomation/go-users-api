package users

// data-access-object - access layer to the database. Persist and retrieve user from database.
// this file is the only place in the app that can access the db.
// this pattern allows for easy management and switching between persistant data bases.

import (
	"github.com/sdetAutomation/go-users-api/datasources/mysql/usersdb"
	"github.com/sdetAutomation/go-users-api/utils/date"
	"github.com/sdetAutomation/go-users-api/utils/errors"
	"github.com/sdetAutomation/go-users-api/utils/sqlhelper"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
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
		return sqlhelper.ParseError(getErr)
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
		return sqlhelper.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return sqlhelper.ParseError(err)
	}
	// update user with last insert id.
	user.ID = userID
	return nil
}

// Update ...
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return sqlhelper.ParseError(err)
	}
	return nil
}
