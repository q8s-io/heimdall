package ginext

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func BodyIntercept() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		// Request.Body for context
		requestBody := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = requestBody
		c.Next()
	}
}
