package image

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/controller"
)

func Routes(router *gin.Engine) {
	images := router.Group("/api/images")
	{
		// image name
		images.POST("/", controller.GetImageInfoByName)
	}
}
