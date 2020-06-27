package ginext

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ResolveJSON(c *gin.Context, reqMap interface{}) {
	err := c.BindJSON(&reqMap)
	if err != nil {
		log.Println("requests body not valid")
	}
}
