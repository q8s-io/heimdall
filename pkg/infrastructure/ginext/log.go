package ginext

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func InitLog() {
	log.SetOutput(os.Stdout)
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStart := time.Now()
		c.Next()
		timeEnd := time.Now()
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		latency := timeEnd.Sub(timeStart)
		comment := c.Errors
		log.Println(clientIP, statusCode, method, path, latency, comment)
	}
}
