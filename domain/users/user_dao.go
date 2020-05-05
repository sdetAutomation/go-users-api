package users

// data-access-object - access layer to the database. Persist and retrieve user from database.
// this file is the only place in the app that can access the db.
// this pattern allows for easy management and switching between persistant data bases.

import (
	"fmt"
	"github.com/sdetAutomation/go-users-api/datasources/mysql/usersdb"
	"github.com/sdetAutomation/go-users-api/utils/errors"
	"github.com/sdetAutomation/go-users-api/utils/sqlhelper"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Get ...
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DataCreated, &user.Status); getErr != nil {
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DataCreated, user.Status, user.Password)
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

// Delete ...
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return sqlhelper.ParseError(err)
	}
	return nil
}

// FindByStatus ...
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DataCreated, &user.Status); err != nil {
			return nil, sqlhelper.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
} 