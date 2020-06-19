package _default

import (
	"github.com/gin-gonic/gin"
	
	"github.com/q8s-io/heimdall/pkg/controller"
)

func Routes(router *gin.Engine) {
	status := router.Group("/api/system")
	{
		// demo
		status.GET("/status", controller.Status)
	}
}
