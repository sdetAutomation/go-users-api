package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sdetAutomation/go-users-api/logger"
)

var(
	router = gin.Default()
)

// StartApplication ...
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8080")
}