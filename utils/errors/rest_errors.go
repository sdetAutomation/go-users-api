package errors

import (
	"net/http"
)

// RestErr ...
type RestErr struct {
	Message	string	`json:"message"`
	Status	int 	`json:"status"`
	Error	string	`json:"error"`
}

// NewBadRequestError ...
func NewBadRequestError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusBadRequest,
		Error: "bad_request",
	}
}

// NewNotFoundError ...
func NewNotFoundError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusNotFound,
		Error: "not_found",
	}
}