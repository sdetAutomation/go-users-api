package app

import "github.com/sdetAutomation/go-users-api/controllers"

func mapUrls() {
	router.GET("/health", controllers.Health)

	router.GET("/users", controllers.GetUsers)

	router.GET("/users/:user_id", controllers.SearchUser)
	
	router.POST("/users", controllers.CreateUser)

	router.PUT("/users/:user_id", controllers.UpdateUser)

	router.DELETE("/users/:user_id", controllers.DeleteUser)
}