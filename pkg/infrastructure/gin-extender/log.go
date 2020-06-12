package gin_extender

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// InitLog is init log info
func InitLog() {
	log.SetOutput(os.Stdout)
}

// GinLogger is gin log format
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
		log.Println(timeEnd.Format("[2006/01/02 15:04:05]"), clientIP, statusCode, method, path, latency, comment)
	}
}

// CustomLogger is custom log for self
func CustomLogger(args ...string) {
	logInfo := strings.Join(args, " ")
	timeMark := time.Now()
	log.Println(timeMark.Format("[2006/01/02 15:04:05]"), logInfo)
}
