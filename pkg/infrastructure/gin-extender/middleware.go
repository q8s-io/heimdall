package gin_extender

import (
	"github.com/gin-gonic/gin"
)

// ResolveJSON is resolve json with BindJSON
func ResolveJSON(c *gin.Context, reqMap interface{}) {
	err := c.BindJSON(&reqMap)
	if err != nil {
		panic("requests body not valid")
	}
}
