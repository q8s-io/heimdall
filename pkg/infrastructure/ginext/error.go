package ginext

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// GinPanic is log panic
func GinPanic(out io.Writer) gin.HandlerFunc {
	data := make([]interface{},0)
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		// rdrc for context
		rdrc := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdrc
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Error Info] %s", err)
				errStack := make([]byte, 1024*16)
				for {
					// get goroutine`s stacktrace, if need all second parameter is true
					size := runtime.Stack(errStack, false)
					// the size of the buffer may be not enough to hold the stacktrace, so double the buffer size
					if size == len(errStack) {
						errStack = make([]byte, len(errStack)<<1)
						continue
					}
					break
				}
				log.Printf("[Panic Info] %s", errStack)
				// return
				Sender(c, 1, "Data format error.", data)
			}
		}()
		c.Next()
	}
}

// ErrorLogger is error log for self
func ErrorLogger(errAgs error) {
	timeMark := time.Now()
	log.Println(timeMark.Format("[2006/01/02 15:04:05]"), errAgs)
}
