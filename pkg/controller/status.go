package controller

import (
	"github.com/gin-gonic/gin"

	ge "github.com/70data/heimdall/pkg/infrastructure/gin-extender"
)

// Status is health check
func Status(c *gin.Context) {
	data := make([]interface{}, 0)
	ge.Sender(c, 0, "This is status.", data)
}
