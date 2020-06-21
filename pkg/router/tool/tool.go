package tool

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/controller"
)

func Routes(router *gin.Engine) {
	tools := router.Group("/api/tools")
	{
		// id
		tools.GET("/id", controller.GetID)
	}
}
