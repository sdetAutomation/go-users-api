package users

// data-access-object - access layer to the database. Persist and retrieve user from database.  
// this file is the only place in the app that can access the db.
// this pattern allows for easy management and switching between persistant data bases.   

import (
	"fmt"

	"github.com/sdetAutomation/go-users-api/datasources/mysql/usersdb"
	"github.com/sdetAutomation/go-users-api/utils/date"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

// Get ...
func (user *User) Get() *errors.RestErr {
	// adding this ping to ensure db connection has been established when app starts. 
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

	result := userDb[user.ID]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DataCreated = result.DataCreated

	return nil
}

// Save ...
func (user *User) Save() *errors.RestErr {
	current := userDb[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	user.DataCreated = date.GetNowString()

	userDb[user.ID] = user

	return nil
}