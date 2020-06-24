package ginext

import (
	"github.com/gin-gonic/gin"
)

func ResolveJSON(c *gin.Context, reqMap interface{}) {
	err := c.BindJSON(&reqMap)
	if err != nil {
		panic("requests body not valid")
	}
}
