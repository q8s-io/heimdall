package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

func Status(c *gin.Context) {
	data := make([]interface{}, 0)
	ginext.Sender(c, 0, "This is status.", data)
}
