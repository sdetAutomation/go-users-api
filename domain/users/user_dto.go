package users

// data-transfer-object - user defined here, and object transferring from persistant layer to our app. 

import (
	"strings"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

// User ...
type User struct {
	ID 			int64	`json:"id"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email 		string	`json:"email"`
	DataCreated string	`json:"date_created"`
}

// // Validate ... this is a function that needs a value to be passed in to exectute.
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address", "please enter an email similar to this format: email@email.com")
// 	}
// 	return nil
// }

// Validate ... this method can check automatically if called without a value being passed in.
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
