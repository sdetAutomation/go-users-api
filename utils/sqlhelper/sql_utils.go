package sqlhelper

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError ...
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequestError("no matching record for given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	// below can tell you error numbers and messages.
	// fmt.Println(sqlErr.Number)
	// fmt.Println(sqlErr.Number)

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
