package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sdetAutomation/go-users-api/domain/users"
	"github.com/sdetAutomation/go-users-api/services"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

func validateUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id, user id should be a number")
	}
	return userID, nil
}

// Welcome ...
func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to my go api!!! go to /users/1 to see what's in the database!!!")
}

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
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		// if there is a database error, return in json format the error.
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// GetUsers ...
func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}

// GetUser ...
func GetUser(c *gin.Context) {
	userID, idErr := validateUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	userID, idErr := validateUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	// unmarshall the json body from the request to the user struct.
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	// Update user record in the datebase
	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		// if there is a database error, return in json format the error.
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	userID, idErr := validateUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusNoContent, nil)

}

// Search ...
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
