package errors

import (
	"net/http"
	
	"github.com/pkg/errors"
)

// RestErr ...
type RestErr struct {
	Message	string	`json:"message"`
	Status	int 	`json:"status"`
	Error	string	`json:"error"`
}

// NewError ...
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError ...
func NewBadRequestError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusBadRequest,
		Error: "bad_request",
	}
}

// NewNotFoundError ...
func NewNotFoundError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusNotFound,
		Error: "not_found",
	}
}

// NewInternalServerError ...
func NewInternalServerError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status:	http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}