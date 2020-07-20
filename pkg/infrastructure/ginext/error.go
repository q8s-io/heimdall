package ginext

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GinPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Error Info] %s", err)
				// return
				data := make([]interface{}, 0)
				Sender(c, 1, "programe error.", data)
			}
		}()
		c.Next()
	}
}
