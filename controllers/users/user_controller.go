package users

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sdetAutomation/go-users-api/domain/users"
	"github.com/sdetAutomation/go-users-api/services"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	var user users.User
	// unmarshall the json body from the request to the user struct.
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// create user record in the datebase
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// if there is a database error, return in json format the error.
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUsers ...
func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}

// GetUser ...
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id, user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}