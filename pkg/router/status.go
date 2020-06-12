package router

import (
	"github.com/gin-gonic/gin"

	"github.com/70data/heimdall/pkg/controller"
)

func statusRoutes(route *gin.Engine) {
	status := route.Group("/api/status")
	{
		// demo
		status.GET("/demo", controller.Status)
	}
}
