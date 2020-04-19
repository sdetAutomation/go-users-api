package users

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdetAutomation/go-users-api/domain/users"
	"github.com/sdetAutomation/go-users-api/services"
	"github.com/sdetAutomation/go-users-api/utils/errors"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	var user users.User

	// *** below code peforms similarly to c.ShouldBindJSON(&user) function below...
	// *** code here for reference. 
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	
	// if err != nil {
	// 	//todo: handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//todo: handle error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// ****

	// unmarshall the json body from the request to the user struct.
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", err.Error())
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

// SearchUser ...
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me please!")
}