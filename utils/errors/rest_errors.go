package errors

import (
	"net/http"
)

// RestErr ...
type RestErr struct {
	Message	string	`json:"message"`
	Status	int 	`json:"status"`
	Error	string	`json:"error"`
	Hint 	string  `json:"hint"`
}

// NewBadRequestError ...
func NewBadRequestError(message string, hint string)*RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusBadRequest,
		Error: "bad_request",
		Hint: hint,
	}
}