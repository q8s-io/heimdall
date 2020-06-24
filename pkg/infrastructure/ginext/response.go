package ginext

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SimpleSender(c *gin.Context, data interface{}) {
	responseJSON, _ := json.Marshal(data)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(200, string(responseJSON))
}

func Sender(c *gin.Context, code int, err string, data interface{}) {
	responseData := map[string]interface{}{
		"code": code,
		"msg":  err,
		"data": data,
	}
	responseJSON, _ := json.Marshal(responseData)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(200, string(responseJSON))
}
