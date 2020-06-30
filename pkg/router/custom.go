package router

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/controller"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

func RouteCustom() *gin.Engine {
	router := gin.New()
	router.Use(ginext.GinLogger())
	router.Use(ginext.Cors())
	router.Use(ginext.GinPanic())

	system := router.Group("/api/system")
	{
		// status
		system.GET("/status", controller.Status)
	}

	return router
}
