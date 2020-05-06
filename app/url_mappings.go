package app

import (
	"github.com/sdetAutomation/go-users-api/controllers/users"
	"github.com/sdetAutomation/go-users-api/controllers/health"
)

func mapUrls() {
	router.GET("/", users.Welcome)

	router.GET("/health", health.Health)
	
	router.GET("/users", users.GetUsers)

	router.GET("/users/:user_id", users.GetUser)
	
	router.POST("/users", users.CreateUser)

	router.PUT("/users/:user_id", users.UpdateUser)

	router.PATCH("/users/:user_id", users.UpdateUser)

	router.GET("/internal/users/search", users.Search)

	router.DELETE("/users/:user_id", users.DeleteUser)
}