package app

import (
	"github.com/sdetAutomation/go-users-api/controllers/users"
	"github.com/sdetAutomation/go-users-api/controllers/health"
)

func mapUrls() {
	router.GET("/health", health.Health)
	
	router.GET("/users", users.GetUsers)

	router.GET("/users/:user_id", users.SearchUser)
	
	router.POST("/users", users.CreateUser)

	router.PUT("/users/:user_id", users.UpdateUser)

	router.DELETE("/users/:user_id", users.DeleteUser)
}