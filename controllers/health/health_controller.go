package health

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Health ...
func Health(c *gin.Context) {
	c.String(http.StatusOK, "I am up, keep on coding the golang!!!")
}