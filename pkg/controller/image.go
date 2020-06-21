package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

// GetImageInfoByName
func GetImageInfoByName(c *gin.Context) {
	body := make(map[string]string)
	ginext.ResolveJSON(c, &body)
	imageName := body["image_name"]
	ginext.CustomLogger(imageName)

	// get result from db

	// trigger pull image

	// return
	data := make(map[string]string)
	ginext.Sender(c, 0, "This is status.", data)
}
