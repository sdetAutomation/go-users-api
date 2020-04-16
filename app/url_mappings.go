package app

import "github.com/sdetAutomation/go-users-api/controllers"

func mapUrls() {
	router.GET("/health", controllers.Health)
}