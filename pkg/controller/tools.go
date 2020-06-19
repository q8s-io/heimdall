package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

// GetID
func GetID(c *gin.Context) {
	// get id
	id := distribution.GetUUID()
	// return
	data := make(map[string]string)
	data["id"] = id
	ginext.Sender(c, 0, "This is status.", data)
}
