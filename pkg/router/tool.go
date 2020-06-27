package router

import (
	"github.com/q8s-io/heimdall/pkg/controller"
)

func RunTool() {
	router := RouteCustom()

	tools := router.Group("/api/tools")
	{
		// get uuid
		tools.GET("/id", controller.GetID)
	}

	_ = router.Run(":12001")
}
