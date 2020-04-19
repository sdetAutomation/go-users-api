package services

import (
	"github.com/sdetAutomation/go-users-api/domain/users"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

// CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}